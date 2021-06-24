package app

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/index"
	"sync"
)

// IndexingService сервис индексации входящих запросов с текстами, используется в Http
type IndexingService struct {
	// настройки
	args *TextAnalyzerArgs
	// текущий индекс статистики
	stats *index.TermStatCollection
	// синхронизатор доступа к статистике и у нас есть именно блокировки на чтение и запись
	statSync sync.RWMutex
}

// NewIndexService создать и инициализировать новый сервис индексации
func NewIndexService(args *TextAnalyzerArgs) *IndexingService {
	result := &IndexingService{args: args}
	result.Reset()
	return result
}

// Reset сброс состояния текущего индекса
func (t *IndexingService) Reset() {
	t.statSync.Lock()
	defer t.statSync.Unlock()
	t.stats = index.NewTermStatCollectionF(t.args.GetStatisticsFilter())
}

// Index индексация переданного текста в текущей коллекции
func (t *IndexingService) Index(part int, text string) {
	// формируем индекс отдельного участка
	subindex := index.CollectFromString(text, index.CollectConfig{Part: part, Filter: t.args.GetStatisticsFilter()})
	// далее синхронно редуцируем результат в состав общей статистики
	t.statSync.Lock()
	defer t.statSync.Unlock()
	t.stats.Merge(subindex)
	t.stats.RebuildFrequencyIndex()
}

// Find возврат top size элементов из индекса
func (t *IndexingService) Find(size int, filter *index.TermFilter) []index.TermStat {
	// тут мы блокируем только на чтение, соответственно разрешает конкурентный доступ
	t.statSync.RLock()
	defer t.statSync.RUnlock()
	referencedResult := t.stats.Find(size, filter)
	// получим копию дереференсированных данных во избежание порчи после передачи
	result := make([]index.TermStat, len(referencedResult))
	for i := 0; i < len(result); i++ {
		result[i] = *referencedResult[i]
	}
	return result
}

package index

import (
	"github.com/comdiv/golang_course_comdiv/internal/textanalyzer/app"
	"sync"
)

// IndexingService сервис индексации входящих запросов с текстами, используется в Http
type IndexingService struct {
	// настройки
	args *app.TextAnalyzerArgs
	// текущий индекс статистики
	stats *TermStatCollection
	// синхронизатор доступа к статистике и у нас есть именно блокировки на чтение и запись
	statSync sync.RWMutex
}

// NewIndexService создать и инициализировать новый сервис индексации
func NewIndexService(args *app.TextAnalyzerArgs) *IndexingService {
	result := &IndexingService{args: args}
	result.Reset()
	return result
}

// Reset сброс состояния текущего индекса
func (t *IndexingService) Reset() {
	t.statSync.Lock()
	defer t.statSync.Unlock()
	t.stats = NewTermStatCollectionF(t.args.GetStatisticsFilter())
}

// Index индексация переданного текста в текущей коллекции
func (t *IndexingService) Index(part int, text string) {
	// формируем индекс отдельного участка
	subindex := CollectFromString(text, CollectConfig{Part: part, Filter: t.args.GetStatisticsFilter()})
	// далее синхронно редуцируем результат в состав общей статистики
	t.statSync.Lock()
	defer t.statSync.Unlock()
	t.stats.Merge(subindex)
}

// Find возврат top size элементов из индекса
func (t *IndexingService) Find(size int, filter *TermFilter) []TermStat {
	// тут мы блокируем только на чтение, соответственно разрешает конкурентный доступ
	t.statSync.RLock()
	defer t.statSync.RUnlock()
	referencedResult := t.stats.Find(size, filter)
	// получим копию дереференсированных данных во избежание порчи после передачи
	result := make([]TermStat,len(referencedResult))
	for i:=0;i<len(result);i++{
		result[i] = *referencedResult[i]
	}
	return result
}

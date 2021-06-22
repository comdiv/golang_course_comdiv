run_lesson_001_hellogo:
	echo Executing lesson #1 `Hello, go!`
	go run cmd/lesson_001_hellogo.go

runall: run_lesson_001_hellogo

test_experiments :
	echo "Testing of different experiments"
	go test ./internal/experiments

test_sorted_linked_list :
	echo "Testing sorted linked list"
	go test github.com/comdiv/golang_course_comdiv/internal/sortedintlist/...
	echo "Benchmarking sorted linked list"
	go test -bench=. -benchmem github.com/comdiv/golang_course_comdiv/internal/sortedintlist/...


test: test_experiments test_sorted_linked_list

taver = 1.6
ta_profile :
	## только не понял как SVG консольно сохранять в нужный файл
	echo "Enter textanalyzer dir"
	cd ./internal/textanalyzer && \
		echo "Execute benchmarks" && \
		# go test --bench=.* --benchmem ./... > benchs_v$(taver).txt && \
		echo "Set last benchmarks as current" && \
		# cp benchs_v$(taver).txt benchs.txt && \
		echo "Collect profiler" && \
		#cat  testdata/large_test_json.json | go run main.go --json --cpuprof cpu.prof --memprof mem.prof && \
		go tool pprof -text cpu.prof > ./pprof/v$(taver)/cpu.txt && \
		go tool pprof -alloc_space -text mem.prof > ./pprof/v$(taver)/alloc_space.txt && \
		go tool pprof -alloc_objects -text mem.prof > ./pprof/v$(taver)/alloc_objects.txt && \
		go tool pprof -svg cpu.prof > ./pprof/v$(taver)/cpu.svg && \
        go tool pprof -alloc_space -svg mem.prof > ./pprof/v$(taver)/alloc_space.svg && \
        go tool pprof -alloc_objects -svg mem.prof > ./pprof/v$(taver)/alloc_objects.svg && \
		cp -rf ./pprof/v$(taver)/* ./pprof/current

ta_http :
	echo "Start text stat service"
	cd ./internal/textanalyzer && \
		go run main.go --first --last --http 8080 --pprofhttp same


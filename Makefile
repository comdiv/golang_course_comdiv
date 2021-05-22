run_lesson_001_hellogo:
	echo Executing lesson #1 `Hello, go!`
	go run cmd/lesson_001_hellogo.go

runall: run_lesson_001_hellogo

test_experiments :
	echo "Testing of different experiments"
	go test ./internal/experiments

test_sorted_linked_list :
	echo "Testing sorted linked list"
	go test github.com/comdiv/golang_course_comdiv/internal/sortedintlist/test
	echo "Benchmarking sorted linked list"
	go test -bench=. -benchmem github.com/comdiv/golang_course_comdiv/internal/sortedintlist/test


test: test_experiments test_sorted_linked_list

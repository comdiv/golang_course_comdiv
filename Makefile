run_lesson_001_hellogo:
	echo Executing lesson #1 `Hello, go!`
	go run cmd/lesson_001_hellogo.go

runall: run_lesson_001_hellogo

test_experiments :
	echo "Testing of different experiments"
	go test ./internal/experiments

test: test_experiments


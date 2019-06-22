default :
	go build main.go
	rm -rf plot1*
	./main 3 0.000001 1 > plot1
	./main 3 0.000001 10 > plot10
	./main 3 0.000001 100 > plot100
	./main 3 0.000001 1000 > plot1000
	./main 3 0.000001 10000 > plot10000
	gnuplot plot.txt
	eog res.png


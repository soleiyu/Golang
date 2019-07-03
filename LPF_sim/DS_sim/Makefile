default :
	go build main.go
	rm -rf pulse*
	./main 0.0001  > pulse1
	gnuplot plot.txt
	eog res.png &


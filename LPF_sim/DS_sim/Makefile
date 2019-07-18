default :
	go build main.go
	rm -rf pulse*
	./main 0.000001  > pulse1
	./main 0.000002  > pulse2
	./main 0.000005  > pulse3
	./main 0.00001  > pulse4
	./main 0.00002  > pulse5
	./main 0.00005  > pulse6
	./main 0.0001  > pulse7
	./main 0.0002  > pulse8
	./main 0.0005  > pulse9
	./main 0.001  > pulse10
	./main 0.002  > pulse11
	./main 0.005  > pulse12
	gnuplot plot.txt
	eog res.png &


default :
	go build main.go
	rm -rf pulse*
	./main 0 > pulse0
	./main 1 > pulse1
	./main 2 > pulse2
	./main 3 > pulse3
	gnuplot plot.txt
	eog res.png &

full :
	go build main.go
	rm -rf pulse*
	./main 100 > pulse
	gnuplot plotf.txt
	eog res.png &

f4 :
	go build main.go
	rm -rf pulse*
	./main 100 0.00000001  > pulse1
	./main 100 0.0000001 > pulse10
	./main 100 0.000001 > pulse100
	./main 100 0.00001 > pulse1000
	gnuplot plotf4.txt
	eog res.png &


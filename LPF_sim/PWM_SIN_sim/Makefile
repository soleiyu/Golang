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
	./main 0 > pulse
	gnuplot plotf.txt
	eog res.png &


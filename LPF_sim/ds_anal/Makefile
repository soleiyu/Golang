default :
	go build main.go
	rm -rf pana*
	./main > pana
	gnuplot plotana.txt
	eog res.png &


default :
	go build main.go
	rm -rf pulse:
	./main > pulse
	gnuplot plot.txt
	eog res.png &


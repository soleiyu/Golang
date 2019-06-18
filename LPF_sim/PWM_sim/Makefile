default :
	go build main.go
	rm -rf plot1*
	./main > pulse
	gnuplot plot.txt
	eog res.png &


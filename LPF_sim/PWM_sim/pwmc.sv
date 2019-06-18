module pwmc (
	input wire clk,
	input wire [9:0] vq,
	output wire pulse
);

reg [9:0] cnt;
wire [10:0] qc = {1'b1, vq} - {1'b0, cnt};

assign pulse = qc[10];

always @(posedge clk) begin
	cnt <= cnt + 10'b1;
end

initial begin
	cnt = 10'b0;
end

endmodule

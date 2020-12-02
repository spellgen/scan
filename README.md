# scan
Define line scanners and read a full heterogeneous file in one go

The format of a particular type of line is defined by an object that implements the LineScanner interface.

    type twoInts struct {
        a int
        b int
    }
    
    func (o twoInts) Parse(s string) (scan.LineScanner, bool) {
    	n, err := fmt.Sscanf(s, "%d,%d\n", &o.a, &o.b)
    	return o, n == 2 && err == nil
    }
    
If the supplied line can be satisfactorily parsed a copy of it is returned.

See the test function for how to use.
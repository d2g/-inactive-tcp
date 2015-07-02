tcp
===

Go TCP packet structure

I have since moved to use: github.com/google/gopacket

    p := gopacket.NewPacket(packet, layers.LayerTypeTCP, gopacket.DecodeOptions{Lazy: true, NoCopy: true})

    if tcpLayer := p.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		    tcp, ok := tcpLayer.(*layers.TCP)
		    if !ok {
			    return ap, errors.New("Unable to convert packet to TCP")
		    }
		
		    //tcp <- has what we want.
    }

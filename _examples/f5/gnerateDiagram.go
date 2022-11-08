package main

import (
	"fmt"
	"log"

	"github.com/jacostaperu/go-diagrams/diagram"
	"github.com/jacostaperu/go-diagrams/nodes/f5bigip"
	"github.com/jacostaperu/go-diagrams/nodes/generic"
)

func URL(c string) diagram.GroupOption {
	return func(o *diagram.GroupOptions) {
		o.Attributes = map[string]string{"URL": "https://google.com"}
	}
}

func Color(c string) diagram.GroupOption {
	return func(o *diagram.GroupOptions) {
		o.Attributes = map[string]string{"fontcolor": "red"}
	}
}
func Tooltip(c string) diagram.GroupOption {
	return func(o *diagram.GroupOptions) {
		o.Attributes = map[string]string{"tooltip": "mytooltip"}
	}
}

func generateDiagram() {
	d, err := diagram.New(diagram.Filename("Radius"), diagram.Label("Radius"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	//dns := gcp.Network.Dns(diagram.NodeLabel("DNS"))
	lb := f5bigip.F5.BigIp(diagram.NodeLabel("f5 slough"))
	d.Add(lb)
	radius1 := generic.Daemon.FreeRadius(diagram.NodeLabel("radius1"))
	radius2 := generic.Daemon.FreeRadius(diagram.NodeLabel("radius2"))
	radius3 := generic.Daemon.FreeRadius(diagram.NodeLabel("radius3"))
	radius4 := generic.Daemon.FreeRadius(diagram.NodeLabel("radius4"))

	dc := diagram.NewGroup("GCP", Color("red"), URL("ddsd"))
	dc.NewGroup("services", Color("red")).
		Label("Service Layer").
		Add(
			radius1,
			radius2,
			radius3,
			radius4,
		)
		//.
		//ConnectAllFrom(lb.ID(), diagram.Forward())
		//.ConnectAllTo(cache.ID(), diagram.Forward())

	dc.NewGroup("ESX1").Label("ESX1").Add(radius1, radius2).
		ConnectAllFrom(lb.ID(), diagram.Forward())

	dc.NewGroup("ESX2").Label("ESX2").Add(radius3, radius4).ConnectAllFrom(lb.ID())
	//d.Connect(dns, lb, diagram.Forward()).Group(dc)
	d.Group(dc)

	dot_as_string, err := d.RenderString()

	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(dot_as_string)
}

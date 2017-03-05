package alg

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "hackupc2017w/carthumbing_api/arcgis"
)

func TestArcGisApi() {
    url := `http://sampleserver3.arcgisonline.com/arcgis/rest/services/Geometry/GeometryServer/intersect?sr=4269&geometries=%7B%22geometryType%22+%3A+%22esriGeometryPolygon%22%2C%22spatialReference%22+%3A%7B%22wkid%22+%3A+4269%7D%2C%22geometries%22+%3A%5B%7B%22rings%22+%3A+%5B%5B%5B-75.48928066099995%2C39.714858219000064%5D%2C%5B-75.4759742679999%2C39.720084384000074%5D%2C%5B-75.47476845699993%2C39.741832093000085%5D%2C%5B-75.46039411899994%2C39.763362027000085%5D%2C%5B-74.73882472699995%2C40.17772564400008%5D%2C%5B-74.9166543419999%2C39.17063854200006%5D%2C%5B-75.01440707699993%2C39.198363837000045%5D%2C%5B-75.11995811199995%2C39.18469178100008%5D%2C%5B-75.4156722749999%2C39.374971842000036%5D%2C%5B-75.55276303999995%2C39.49051430700007%5D%2C%5B-75.5166888839999%2C39.56656841600005%5D%2C%5B-75.57023418699993%2C39.61773496300009%5D%2C%5B-75.48928066099995%2C39.714858219000064%5D%5D%5D%7D%5D%7D&geometry=%7B%22geometryType%22+%3A+%22esriGeometryPolygon%22%2C%22spatialReference%22+%3A%7B%22wkid%22+%3A+4269%7D%2C%22geometry%22+%3A%7B%22rings%22+%3A+%5B%5B%5B-75.48928066099995%2C39.714858219000064%5D%2C%5B-75.4759742679999%2C39.720084384000074%5D%2C%5B-75.47476845699993%2C39.741832093000085%5D%2C%5B-75.46039411899994%2C39.763362027000085%5D%2C%5B-74.73882472699995%2C40.17772564400008%5D%2C%5B-75.48928066099995%2C39.714858219000064%5D%5D%5D%7D%7D&f=pjson`
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    s := string(body[:])
    fmt.Println(s)
}


func TestRoute() {
    p1 := arcgis.Point{-122.4079,37.78356}
    p2 := arcgis.Point{-122.404,37.782}
    // points, err := arcgis.GetRoutePath(&p1, &p2)
    // if err != nil {
    //     panic(err)
    // }
    // for _, point := range points {
    //     fmt.Println(point)
    // }
    geom, err := arcgis.GetRouteGeometry(&p1, &p2)
    if err != nil {
        panic(err)
    }
    _, err = arcgis.BufferPathGeom(geom)
    if err != nil {
        panic(err)
    }
}
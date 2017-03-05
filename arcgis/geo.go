package arcgis

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "bytes"
)

type GeometryPoint struct {
    Coords []float64
}

type Geometry struct {
    Points [][][]float64 `json:"paths"`
}

type GeometryObject struct {
    GeometryType string `json:"geometryType"`
    Geometries []Geometry `json:"geometries"`
}

func getBufferData(path []*Point) ([]byte, error) {
    obj := GeometryObject{}
    obj.GeometryType = "esriGeometryPolyline"
    obj.Geometries = make([]Geometry, 1)
    obj.Geometries[0] = Geometry{}
    obj.Geometries[0].Points[0] = make([][]float64, 1)
    for _, p := range path {
        gp := []float64{p.Lat, p.Lon}
        obj.Geometries[0].Points[0] = append(obj.Geometries[0].Points[0], gp)
    }
    fmt.Println(obj)
    b, err := json.Marshal(obj)
    if err != nil {
        return nil, err
    }
    s := string(b[:])
    fmt.Println(s)
    return b, nil
}

func BufferPath(path []*Point) (*Polygon, error){
    b, err := getBufferData(path)
    if err != nil {
        return nil, err
    }
    url := "http://sampleserver3.arcgisonline.com/arcgis/rest/services/Geometry/GeometryServer/buffer?&inSR=4326&outSR=4326&bufferSR=4326&distances=1000&unit=&unionResults=true&f=pjson"
    url = url + "&geometries=" + string(b[:])
    req, err := http.NewRequest("GET", url, bytes.NewBuffer(b))
    if err != nil {
        return nil, err

    }
    fmt.Println(url)
    //return nil, nil
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
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
    return nil, nil
}

func BufferPathGeom(geom interface{}) (*Polygon, error){
    d := map[string]interface{}{
        "geometryType": "esriGeometryPolyline",
        "geometries": geom,
    }
    b, err := json.Marshal(d)
    if err != nil {
        return nil, err
    }
    url := "http://sampleserver3.arcgisonline.com/arcgis/rest/services/Geometry/GeometryServer/buffer?&inSR=4326&outSR=4326&bufferSR=4326&distances=10000&unit=&unionResults=false&f=pjson"
    url = url + "&geometries=" + string(b[:])
    fmt.Println(url)
    req, err := http.NewRequest("GET", url, bytes.NewBuffer(b))
    if err != nil {
        return nil, err
    }
    //return nil, nil
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
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
    return nil, nil
}

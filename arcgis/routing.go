package arcgis

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
)

func getRouteResponse(origin, dest *Point) ([]byte, error) {
    // TODO: temporal API key for test purposes
    key := "g_dghjVLRCNdtuSW5iu2E5Qz40kMODyxjJw_nP_S1yTjiSnkswi5zxKxtUAGFd6iYIhC_sBvfgkBv6z18biS" +
           "xFcGS8oOF2AZWgaTO4riurueN9QEL9RGciPRhDuBWXqZPAHm3fwyGVULNccIFiI73A"
    url := "http://route.arcgis.com/arcgis/rest/services/World/Route/NAServer/Route_World/solve?" +
           "token=%s&stops=%f,%f;%f,%f&f=json"
    url = fmt.Sprintf(url, key, origin.Lat, origin.Lon, dest.Lat, dest.Lon)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    s := string(body[:])
    fmt.Println(s)
    return body, nil
}

func GetRoutePath(origin, dest *Point) ([]*Point, error){
    resp, err := getRouteResponse(origin, dest)
    if err != nil {
        return nil, err
    }
    // TODO: use proper structs for unmarshaling
    var data interface{}
    json.Unmarshal(resp, &data)
    r := data.(map[string]interface{})["routes"]
    f := r.(map[string]interface{})["features"]
    f0 := f.([]interface{})[0]
    g := f0.(map[string]interface{})["geometry"]
    p := g.(map[string]interface{})["paths"]
    pd := p.([]interface{})[0]

    points := make([]*Point, 0)
    for _, latlong := range pd.([]interface{}) {
        lat := latlong.([]interface{})[0]
        long := latlong.([]interface{})[1]
        point := Point{lat.(float64), long.(float64)}
        points = append(points, &point)

    }
    return points, nil
}


func GetRouteGeometry(origin, dest *Point) (interface{}, error){
    resp, err := getRouteResponse(origin, dest)
    if err != nil {
        return nil, err
    }
    // TODO: use proper structs for unmarshaling
    var data interface{}
    json.Unmarshal(resp, &data)
    r := data.(map[string]interface{})["routes"]
    f := r.(map[string]interface{})["features"]
    f0 := f.([]interface{})[0]
    g := f0.(map[string]interface{})["geometry"]
    return g, nil
}
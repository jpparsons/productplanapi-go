# productplanapi-go

https://app.productplan.com/docs

### Example client
```go
package main

import (
	"fmt"
	"os"

	"github.com/jpparsons/productplanapi-go/productplan"
)

func main() {
  oauthToken := os.Getenv("TOKEN")
  url := "https://app.productplan.com"

  client := productplan.NewClient(url, productplan.NewOauthTokenCredentials(oauthToken))
  statusResponse, err := client.Status.GetStatus()
  if err != nil {
    fmt.Printf("Status() returned error: %v\n", err)
    os.Exit(1)
  }
  // List Roadmaps
  roadmapList, err := client.Roadmaps.ListRoadmaps(&productplan.ListOptions{Filters: "name=Planning Roadmap*"})
  if err != nil {
    fmt.Printf("ListRoadmaps() returned error: %v\n", err)
    os.Exit(1)
  }

  for _, roadmap := range *roadmapList {
    fmt.Println(roadmap.Name)
    fmt.Println(roadmap.ID) 
  }
  
  // Get bars for a roadmap
  roadmap := (*roadmapList)[0]
  bars, err := client.Roadmaps.GetBars(roadmap.Roadmap)
  if err != nil {
    fmt.Printf("GetBars() returned error: %v\n", err)
    os.Exit(1)
  }

  for _, bar := range *bars {
    fmt.Println(bar.ID)
    fmt.Println(bar.Name)
  }
  
  // Import an idea (parkinglot bar)
  idea := productplan.Ideas{
    Name:           "Docker",
    Description:    "Setup Docker",
    StrategicValue: "High",
    Notes:          "isolated",
    PercentDone:    15,
    Effort:         5,
    Tags:           []string{"devops", "security"},
    Fields:         map[string]string{"pp_lanes": "Lane 2", "pp_legend": "Goal 4"}, 
  }

  ideas := productplan.IdeasImportAttributes{
    IdeaImportRoadmap: productplan.IdeaImportRoadmap{ID: roadmap.ID},
    Ideas:             []productplan.Ideas{idea},
  }

  // import returns no response data (resp will contain the code as resp.HTTPResponse.StatusCode)
  _, err = client.Ideas.Import(ideas)
  if err != nil {
    fmt.Printf("Ideas.Import() returned error: %v\n", err)
  }
}
```

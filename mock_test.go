package mixpanel

import (
	"fmt"
	"time"
)

var fullfillsInterface Mixpanel = &Mock{}

func ExampleMock() {
	client := NewMock()

	t, _ := time.Parse(time.RFC3339, "2016-03-03T15:17:53+01:00")

	client.Update("1", &Update{
		Operation: "$set",
		Timestamp: &t,
		IP:        "127.0.0.1",
		Properties: map[string]interface{}{
			"custom_field": "cool!",
		},
	})

	client.Track("1", "Sign In", &Event{
		IP: "1.2.3.4",
		Properties: map[string]interface{}{
			"from": "email",
		},
	})

	fmt.Println(client)

	// Output:
	// 1:
	//   ip: 127.0.0.1
	//   time: 2016-03-03T15:17:53+01:00
	//   properties:
	//     custom_field: cool!
	//   events:
	//     Sign In:
	//       IP: 1.2.3.4
	//       Timestamp:
	//       from: email
}

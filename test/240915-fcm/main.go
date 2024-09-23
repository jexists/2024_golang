package main

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
	fcm "github.com/appleboy/go-fcm"
)

func main() {
	ctx := context.Background()
	client, err := fcm.NewClient(
		ctx,
		fcm.WithCredentialsFile("mation-test-firebase-adminsdk-nfemk-c96a56fe92.json"),
		// initial with service account
		// fcm.WithServiceAccount("my-client-id@my-project-id.iam.gserviceaccount.com"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Send to a single device
	token := "eyfqFSnmTQOq4ivnFp2X4Y:APA91bHP2D6mJzA0GYKmNkoNJllolN6HAmruiXl5AG44r2FpFdRpw7_OFrpjRhfWWw0dgB_wOVWNSaaFsDUqmvm190S1bMwCT3D7tgSchFmXv8_0Cqarbp1ASfKGeWeO0O0YM0M5Li59"

	messages := []*messaging.Message{
		{
			Token: token,
			Notification: &messaging.Notification{
				Title: "test",
				Body:  "body",
				// ImageURL: os.Getenv("ImageCDN") + notiImageURL, // 이미지는 안보임
			},
			Data: map[string]string{
				"link":          "chat",
				"message_count": fmt.Sprintf("%v", 5),
				"send_username": "jexists",
			},
			Android: &messaging.AndroidConfig{
				Notification: &messaging.AndroidNotification{
					Tag: "brace.formation",
				},
			},
			APNS: &messaging.APNSConfig{
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						MutableContent: true,
						Alert: &messaging.ApsAlert{
							Title: "test",
							Body:  "body",
							// LaunchImage: os.Getenv("ImageCDN") + notiImageURL, // 이미지는 안보임
						},
						Sound: "default",
					},
				},
			},
		},
		// {
		// 	Token: token,
		// 	Notification: &messaging.Notification{
		// 		Title: "test1",
		// 		Body:  "body1",
		// 		// ImageURL: os.Getenv("ImageCDN") + notiImageURL, // 이미지는 안보임
		// 	},
		// 	Data: map[string]string{
		// 		"link":          "chat",
		// 		"message_count": fmt.Sprintf("%v", 5),
		// 		"send_username": "jexists",
		// 	},
		// 	Android: &messaging.AndroidConfig{
		// 		Notification: &messaging.AndroidNotification{
		// 			Tag: "brace.formation",
		// 		},
		// 	},
		// 	APNS: &messaging.APNSConfig{
		// 		Payload: &messaging.APNSPayload{
		// 			Aps: &messaging.Aps{
		// 				MutableContent: true,
		// 				Alert: &messaging.ApsAlert{
		// 					Title: "test1",
		// 					Body:  "body1",
		// 					// LaunchImage: os.Getenv("ImageCDN") + notiImageURL, // 이미지는 안보임
		// 				},
		// 				Sound: "default",
		// 			},
		// 		},
		// 	},
		// },
	}
	resp, err := client.Send(
		ctx,
		messages...,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success count:", resp.SuccessCount)
	fmt.Println("failure count:", resp.FailureCount)
	fmt.Println("message id:", resp.Responses[0].MessageID)
	fmt.Println("error msg:", resp.Responses[0].Error)

}

func send() {
	ctx := context.Background()
	client, err := fcm.NewClient(
		ctx,
		fcm.WithCredentialsFile("mation-test-firebase-adminsdk-nfemk-c96a56fe92.json"),
		// initial with service account
		// fcm.WithServiceAccount("my-client-id@my-project-id.iam.gserviceaccount.com"),
	)
	if err != nil {
		log.Fatal(err)
	}
	token := "eyfqFSnmTQOq4ivnFp2X4Y:APA91bHP2D6mJzA0GYKmNkoNJllolN6HAmruiXl5AG44r2FpFdRpw7_OFrpjRhfWWw0dgB_wOVWNSaaFsDUqmvm190S1bMwCT3D7tgSchFmXv8_0Cqarbp1ASfKGeWeO0O0YM0M5Li59"
	resp, err := client.Send(
		ctx,
		&messaging.Message{
			// Token: token,
			// Data: map[string]string{
			// 	"link":          "chat",
			// 	"message_count": fmt.Sprintf("%v", 5),
			// 	"send_username": "jexists",
			// },

			Token: token,
			Notification: &messaging.Notification{
				Title: "test",
				Body:  "body",
				// ImageURL: os.Getenv("ImageCDN") + notiImageURL, // 이미지는 안보임
			},
			Data: map[string]string{
				"link":          "chat",
				"message_count": fmt.Sprintf("%v", 5),
				"send_username": "jexists",
			},
			Android: &messaging.AndroidConfig{
				Notification: &messaging.AndroidNotification{
					Tag: "brace.formation",
				},
			},
			APNS: &messaging.APNSConfig{
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						MutableContent: true,
						Alert: &messaging.ApsAlert{
							Title: "test",
							Body:  "body",
							// LaunchImage: os.Getenv("ImageCDN") + notiImageURL, // 이미지는 안보임
						},
						Sound: "default",
					},
				},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success count:", resp.SuccessCount)
	fmt.Println("failure count:", resp.FailureCount)
	fmt.Println("message id:", resp.Responses[0].MessageID)
	fmt.Println("error msg:", resp.Responses[0].Error)
}

// // Send to topic
// resp, err = client.Send(
// 	ctx,
// 	&messaging.Message{
// 		Data: map[string]string{
// 			"foo": "bar",
// 		},
// 		Topic: "highScores",
// 	},
// )
// if err != nil {
// 	log.Fatal(err)
// }

// // Send with condition
// // Define a condition which will send to devices which are subscribed
// // to either the Google stock or the tech industry topics.
// condition := "'stock-GOOG' in topics || 'industry-tech' in topics"

// // See documentation on defining a message payload.
// message := &messaging.Message{
// 	Data: map[string]string{
// 		"score": "850",
// 		"time":  "2:45",
// 	},
// 	Condition: condition,
// }

// resp, err = client.Send(
// 	ctx,
// 	message,
// )
// if err != nil {
// 	log.Fatal(err)
// }

// // Send multiple messages to device
// // Create a list containing up to 500 messages.
// registrationToken := "YOUR_REGISTRATION_TOKEN"
// messages := []*messaging.Message{
// 	{
// 		Notification: &messaging.Notification{
// 			Title: "Price drop",
// 			Body:  "5% off all electronics",
// 		},
// 		Token: registrationToken,
// 	},
// 	{
// 		Notification: &messaging.Notification{
// 			Title: "Price drop",
// 			Body:  "2% off all books",
// 		},
// 		Topic: "readers-club",
// 	},
// }
// resp, err = client.Send(
// 	ctx,
// 	messages...,
// )
// if err != nil {
// 	log.Fatal(err)
// }

// // Send multicast message
// // Create a list containing up to 500 registration tokens.
// // This registration tokens come from the client FCM SDKs.
// registrationTokens := []string{
// 	"eyfqFSnmTQOq4ivnFp2X4Y:APA91bHP2D6mJzA0GYKmNkoNJllolN6HAmruiXl5AG44r2FpFdRpw7_OFrpjRhfWWw0dgB_wOVWNSaaFsDUqmvm190S1bMwCT3D7tgSchFmXv8_0Cqarbp1ASfKGeWeO0O0YM0M5Li59",
// 	// ...
// 	"eyfqFSnmTQOq4ivnFp2X4Y:APA91bHP2D6mJzA0GYKmNkoNJllolN6HAmruiXl5AG44r2FpFdRpw7_OFrpjRhfWWw0dgB_wOVWNSaaFsDUqmvm190S1bMwCT3D7tgSchFmXv8_0Cqarbp1ASfKGeWeO0O0YM0M5Li591",
// }
// msg := &messaging.MulticastMessage{
// 	Data: map[string]string{
// 		"score": "850",
// 		"time":  "2:45",
// 	},
// 	Tokens: registrationTokens,
// }
// resp, err = client.SendMulticast(
// 	ctx,
// 	msg,
// )
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Printf("%d messages were sent successfully\n", resp.SuccessCount)
// if resp.FailureCount > 0 {
// 	var failedTokens []string
// 	for idx, resp := range resp.Responses {
// 		if !resp.Success {
// 			// The order of responses corresponds to the order of the registration tokens.
// 			failedTokens = append(failedTokens, registrationTokens[idx])
// 		}
// 	}

// 	fmt.Printf("List of tokens that caused failures: %v\n", failedTokens)
// }

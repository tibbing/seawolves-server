// package util

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"gopkg.in/mgo.v2/bson"
// )

// // Timestamp Json compatible timestamp
// type Timestamp time.Time

// // MarshalJSON function
// func (t *Timestamp) MarshalJSON() ([]byte, error) {
// 	ts := time.Time(*t).Unix()
// 	stamp := fmt.Sprint(ts)

// 	return []byte(stamp), nil
// }

// // UnmarshalJSON function
// func (t *Timestamp) UnmarshalJSON(b []byte) error {
// 	ts, err := strconv.Atoi(string(b))
// 	if err != nil {
// 		return err
// 	}

// 	*t = Timestamp(time.Unix(int64(ts), 0))

// 	return nil
// }

// // GetBSON function
// func (t Timestamp) GetBSON() (interface{}, error) {
// 	if time.Time(t).IsZero() {
// 		return nil, nil
// 	}

// 	return time.Time(t), nil
// }

// // SetBSON function
// func (t *Timestamp) SetBSON(raw bson.Raw) error {
// 	var tm time.Time

// 	if err := raw.Unmarshal(&tm); err != nil {
// 		return err
// 	}

// 	*t = Timestamp(tm)

// 	return nil
// }

// func (t *Timestamp) String() string {
// 	return time.Time(*t).String()
// }

// //*util.Timestamp `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
// //now := time.Now().UTC().UnixNano() // time.Now()
// //(*util.Timestamp)(&now)

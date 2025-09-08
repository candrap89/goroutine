package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func prime(n int, output chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var prime []int
	for i := 0; i < n; i++ {
		if isPrima(i) {
			prime = append(prime, i)
		}
	}
	output <- prime
}

func isPrima(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func fibonaci(n int, ch chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n <= 0 {
		ch <- []int{}
	}

	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	ch <- fib

}

func main() {

	generateJWTCreateOrder()
	generateJWTListOrder()
	generateJWTUpdateOrder()

	// Channel untuk komunikasi antar goroutine
	primeCh := make(chan []int)
	fibCh := make(chan []int)

	// contoh concurency
	var wg sync.WaitGroup
	wg.Add(2)

	// Menjalankan fungsi dalam goroutine
	go prime(10, primeCh, &wg)
	go fibonaci(10, fibCh, &wg)

	// Menunggu goroutine selesai
	go func() {
		wg.Wait()
		close(primeCh)
		close(fibCh)
	}()

	// Menerima hasil dari channel
	primes := <-primeCh
	fibonacci := <-fibCh

	// Menampilkan hasil
	fmt.Println("Bilangan Prima:", primes)
	fmt.Println("Bilangan Fibonacci:", fibonacci)

}

func base64UrlEncode(data []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func signHS256(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	signature := h.Sum(nil)
	return base64UrlEncode(signature)
}

func randomString(n int) string {
	result := make([]byte, n)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		result[i] = letterBytes[num.Int64()]
	}
	return string(result)
}

func randomDigits(n int) string {
	result := make([]byte, n)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(10)) // digits 0-9
		result[i] = byte('0') + byte(num.Int64())
	}
	return string(result)
}

func generateJWTCreateOrder() {
	secret := "123456"
	account := "SHOPEE"
	timestamp := uint(time.Now().Unix())

	// 1. Header (include account)
	header := map[string]interface{}{
		"alg":     "HS256",
		"typ":     "JWT",
		"account": account,
	}
	headerBytes, _ := json.Marshal(header)
	headerEncoded := base64UrlEncode(headerBytes)
	fmt.Println("timestamp : ", timestamp)
	// 2. Payload
	data := map[string]interface{}{
		"idempotency_key": uuid.New().String(),
		//"idempotency_key": "a27563ad-58a5-4cbd-a785-f096f5d35ab5",
		"order": map[string]interface{}{
			"reference_no":               "REF" + randomString(4) + randomDigits(4) + fmt.Sprint(timestamp),
			"business_type":              1,
			"delivery_type":              2,
			"is_reverse":                 0,
			"carrier_tn":                 "",
			"total_weight":               410,
			"estimate_chargeable_weight": 410,
			"goods_value":                "85025",
			"goods_value_non_tax":        "80000",
		},
		"item_info": []map[string]interface{}{
			{
				"item_name":        "Celana Chino Pendek / RIP Pria Kualitas Distro Premium",
				"item_quantity":    1,
				"item_price":       3423,
				"category":         "Men Clothes",
				"sub_category":     "Pants",
				"sub_sub_category": "Chino",
				"height":           0,
				"width":            0,
				"length":           0,
			},
			{
				"item_name":        "Kaos Polos Premium Cotton Combed 30s",
				"item_quantity":    2,
				"item_price":       1500,
				"category":         "Men Clothes",
				"sub_category":     "Shirts",
				"sub_sub_category": "T-Shirts",
				"height":           0,
				"width":            0,
				"length":           0,
			},
		},
		"sender_info": map[string]interface{}{
			"name":           "John Doe",
			"phone":          "123456789",
			"country":        "SG",
			"state":          "Singapore",
			"city":           "Singapore",
			"detail_address": "123 Orchard Road",
			"remark":         "Handle with care",
		},
		"receiver_info": map[string]interface{}{
			"name":           "Jane Smith",
			"phone":          "987654321",
			"country":        "SG",
			"state":          "Singapore",
			"city":           "Singapore",
			"detail_address": "456 Marina Bay",
			"remark":         "Fragile",
		},
		"service_info": map[string]interface{}{
			"cod_amount":                "1000",                               // Example COD amount
			"service_types":             "11,21,22",                           // Example service types
			"schedule_pickup_date_time": time.Now().Add(1 * time.Hour).Unix(), // Example schedule pickup date time
			"schedule_pickup_timeslot": map[string]interface{}{
				"start_time": time.Now().Add(1 * time.Hour).Unix(),
				"end_time":   time.Now().Add(2 * time.Hour).Unix(),
			},
		},
	}

	payload := map[string]interface{}{
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
	payloadBytes, _ := json.Marshal(payload)
	payloadEncoded := base64UrlEncode(payloadBytes)

	// 3. Signature
	signingInput := headerEncoded + "." + payloadEncoded
	signature := signHS256(signingInput, secret)

	// 4. JWT Token
	jwtToken := signingInput + "." + signature

	fmt.Println("Generated JWT Token Create Order:")
	fmt.Println(jwtToken)
}

func generateJWTListOrder() {
	secret := "123456"
	account := "SHOPEE"

	// 1. Header (include account)
	header := map[string]interface{}{
		"alg":     "HS256",
		"typ":     "JWT",
		"account": account,
	}
	headerBytes, _ := json.Marshal(header)
	headerEncoded := base64UrlEncode(headerBytes)

	// 2. Payload
	data := map[string]interface{}{
		"pageno": 2, // starting from 0,
		"count":  5, // limit to 10 items
	}

	payload := map[string]interface{}{
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
	payloadBytes, _ := json.Marshal(payload)
	payloadEncoded := base64UrlEncode(payloadBytes)

	// 3. Signature
	signingInput := headerEncoded + "." + payloadEncoded
	signature := signHS256(signingInput, secret)

	// 4. JWT Token
	jwtToken := signingInput + "." + signature

	fmt.Println("Generated JWT List Order:")
	fmt.Println(jwtToken)
}

func generateJWTUpdateOrder() {
	secret := "123456"
	account := "SHOPEE"

	// 1. Header (include account)
	header := map[string]interface{}{
		"alg":     "HS256",
		"typ":     "JWT",
		"account": account,
	}
	headerBytes, _ := json.Marshal(header)
	headerEncoded := base64UrlEncode(headerBytes)

	// 2. Payload
	data := map[string]interface{}{
		"idempotency_key": uuid.New().String(),
		"carrier_tn":      "CTNQain2223",
		"sender_info": map[string]interface{}{
			"name":           "candra",
			"phone":          "123456789",
			"country":        "SG",
			"state":          "Singapore",
			"city":           "Singapore",
			"detail_address": "123 Orchard Road",
			"remark":         "Handle with care",
		},
		"receiver_info": map[string]interface{}{
			"name":           "Jane Smith",
			"phone":          "987654321",
			"country":        "SG",
			"state":          "Singapore",
			"city":           "Singapore",
			"detail_address": "456 Marina Bay",
			"remark":         "Fragile",
		},
		"service_info": map[string]interface{}{
			"service_types": "Express",
			"schedule_pickup_timeslot": map[string]interface{}{
				"start_time": time.Now().Add(1 * time.Hour).Unix(),
				"end_time":   time.Now().Add(2 * time.Hour).Unix(),
			},
		},
	}

	payload := map[string]interface{}{
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
	payloadBytes, _ := json.Marshal(payload)
	payloadEncoded := base64UrlEncode(payloadBytes)

	// 3. Signature
	signingInput := headerEncoded + "." + payloadEncoded
	signature := signHS256(signingInput, secret)

	// 4. JWT Token
	jwtToken := signingInput + "." + signature

	fmt.Println("Generated JWT Token Update Order:")
	fmt.Println(jwtToken)
}

// json example
// {
//   "order": {
//     "reference_no": "REF123456789",
//     "business_type": 1,
//     "delivery_type": 2,
//     "is_reverse": 0,
//     "carrier_tn": "",
//     "total_weight": 410,
//     "estimate_chargeable_weight": 410,
//     "goods_value": "85025",
//     "goods_value_non_tax": "80000"
//   },
//   "item_info": [
//     {
//       "item_name": "Celana Chino Pendek / RIP Pria Kualitas Distro Premium",
//       "item_quantity": 1,
//       "item_price": 3423,
//       "category": "Men Clothes",
//       "sub_category": "Pants",
//       "sub_sub_category": "Chino",
//       "height": 0,
//       "width": 0,
//       "length": 0
//     },
//     {
//       "item_name": "Kaos Polos Premium Cotton Combed 30s",
//       "item_quantity": 2,
//       "item_price": 1500,
//       "category": "Men Clothes",
//       "sub_category": "Shirts",
//       "sub_sub_category": "T-Shirts",
//       "height": 0,
//       "width": 0,
//       "length": 0
//     }
//   ],
//   "sender_info": {
//     "name": "John Doe",
//     "phone": "123456789",
//     "country": "SG",
//     "state": "Singapore",
//     "city": "Singapore",
//     "detail_address": "123 Orchard Road",
//     "remark": "Handle with care"
//   },
//   "receiver_info": {
//     "name": "Jane Smith",
//     "phone": "987654321",
//     "country": "SG",
//     "state": "Singapore",
//     "city": "Singapore",
//     "detail_address": "456 Marina Bay",
//     "remark": "Fragile"
//   },
//   "service_info": {
//     "service_types": "Express",
//     "schedule_pickup_timeslot": {
//       "start_time": 1749493826,
//       "end_time": 1749497426
//     }
//   }
// }

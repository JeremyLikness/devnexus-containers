package main

import (
	"math/rand"
    "fmt"
    "net/http"
	"time"
)

var serverID = rand.New(
	rand.NewSource(
		time.Now().UnixNano())).Intn(3)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var fn [3]string

	// credit: https://www.dwitter.net/d/1978
	fn[0] = "for(d=256;--d;)for(X=-32;++X<32;x.fillRect(960+X*8e3/d,100+7e4/d-h,120,9))x.fillStyle=R(h=(C(a=d/9+t*6)+C(X/9+a/3))*35,g=h+C(X^d)*6+90,g/.8)"
	
	// credit: https://www.dwitter.net/d/4388
	fn[1] = "for(i=w=3e3,a=b=g=333*T(t%9/7);j=(i^a*a)%4,i--;x.fillRect(a+=g+C(j*5+t)*g-a/2-w+g,b+=g+S(j*9+t)*g-b/2-w,s,s),w=2)x.fillStyle=R(s=w*9,w,s,.1)"
	
	// credit: https://www.dwitter.net/d/4136
	fn[2] = "t/=4;for(i=0;i<1100;i++){d=C(t+i),s=i==0?2920:9-d*8;x.fillStyle=R(i,i,i,0.1);x.fillRect(S(t*2.5+i)*220*S(C(t)+i%50)+960*(i==0?-1:1),i,s,s);}"
	
	fmt.Fprintln(w, fn[serverID], r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
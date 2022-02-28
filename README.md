]# Intro
- [Naver 지도 API](https://www.ncloud.com/product/applicationService/maps) 를 기반으로 
지도상의 영역을 100m^2 넓이의 셀로 나누어 DB에 저장합니다.
- 좌표계는 네이버 지도에서 사용하는 `UTM-K` 좌표계를 사용합니다.
    - 좌표계 재택이유는 `UTM-K` 좌표계가 타원체가 아닌 평면체에 투영하여 지도를 보여주기 때문에 평면 View에서 가장 오차가 적음

# Project Structure
```
├── README.md
├── cell                        지도상에 그려질 Cell 객체 관리
│   ├── cell.go                 
│   ├── cell_test.go           
│   └── crud.go                 
├── const                       각종 설정값 (API 키/DB 연결 정보) 상수
│   └── constlocal.go           
├── db                          Database ORM 관련 코드
│   ├── db.go                   
│   └── db_test.go              
├── go.mod                      
├── go.sum                      
├── main.go                     메인 함수
├── naverapi                    네이버 API 통신을 위한 코드
│   ├── naverapi.go             
│   └── naverapi_test.go        
└── utils                       각종 유틸 함수 모음 
    └── utils.go
```


# Usage
```bash
// const/constlocal.go

package _const

const (
NCPKEYID = ""       // 네이버 API Client ID
NCPKEY = ""         // 네이버 API Client Secret
MYSQLUSER = ""      // MySQL DB 접속 id
MYSQLPASS = ""      // MySQL DB 접속 pw
MYSQLHOST = ""      // MySQL DB 호스트
MYSQLDB = ""        // MySQL DB 이름
)

```

```bash
// Test

go test ./...
>>>>
?       handlegeo       [no test files]
ok      handlegeo/cell  1.052s
?       handlegeo/const [no test files]
ok      handlegeo/db    0.990s
ok      handlegeo/naverapi      0.741s
?       handlegeo/utils [no test files]
```

```bash
// run

go run main.go
```


# TODO.
- REST API 혹은 CLI 로 컨트롤할 수 있는 환경 만들기
- 코드 주석 한글로 달기
- 입력된 셀 관리 기능 추가

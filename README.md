# woowacourse-openmission

## 참고 자료
- MSA에 대해 관심을 갖게됨
  - https://www.youtube.com/watch?v=BnS6343GTkY&t=1900s
- MSA와 kafka에 대해 찾아봄
  - https://www.youtube.com/playlist?list=PLtUgHNmvcs6p3304vUg6ywvUIAIe5K4WP
  - https://www.youtube.com/playlist?list=PLtUgHNmvcs6qOCmm-RRK-c5CEMjW_PGlK
- MSA에서 Go를 도입계기를 알게됨
  - https://www.youtube.com/watch?v=mLIthm96u2Q
- Go의 프로그래밍 가이드 번역된 공식문서 읽기 도전
  - https://gosudaweb.gitbooks.io/effective-go-in-korean/content/
- Go 문법에 대해 학습
  - https://www.youtube.com/watch?v=8uiZC0l4Ajw&t=1s
- Go 문법 공식문서와 Gin 공식문서 읽기 도전
  - https://go.dev/ref/spec
  - https://gin-gonic.com/en/docs/
- Gin 프로젝트 구조에 대해 공부
  - https://velog.io/@fudoge/Gin-Gin-%ED%94%84%EB%A1%9C%EC%A0%9D%ED%8A%B8-%EA%B5%AC%EC%A1%B0-%EC%84%A4%EC%A0%95

## 미션 내용
- 달러 기준 쿠폰과 결제할인을 고려한 최저가 찾아주기
- 결제는 비동기 처리로 진행
- 배포하기

## 내용 선택 계기
- 금액별 할인율이 일정하지 않아 결제전 최저가를 계산하는데 너무 오래걸림

### 입력
- 구매할 상품의 개수 입력
- 구매할 상품의 이름과 가격 입력
- 결제 여부 입력


### 출력
- 현재 사용가능한 쿠폰과 결제할인 출력
- 상품 그룹별 할인 전 가격과 할인 후 가격 출력 
- 현재 환율 출력
- 결제 성공여부 출력


```
쿠폰 할인 정보
25달러 이상 구매하면 4달러 할인
54달러 이상 구매하면 8달러 할인
89달러 이상 구매하면 13달러 할인
209달러 이상 구매하면 31달러 할인
679달러 이상 구매하면 102달러 할인

카드 할인 정보
50달러 이상 구매하면 8달러 할인
100달러 이상 구매하면 16달러 할인
150달러 이상 구매하면 30달러 할인
300달러 이상 구매하면 60달러 할인

구매하고 싶은 상품의 개수를 입력해주세요.
8

구매하고 싶은 상품의 이름과 가격(달러)을 입력해주세요.
A 10
B 20
C 30
D 40
E 10
F 20
G 30
H 40

A, E, F, H (Before: 80, After: 64)
B, D (Before: 60, After: 44)
C, G (Before: 60, After: 44)

현재 환율 정보
1USD = 1,483.08원

결제 하시러면 y를 입력해주세요
y

A, E, F, H : 64달러 결제 성공

B, D : 44달러 결제 성공

C, G : 44달러 결제 성공

프로그램 종료
```

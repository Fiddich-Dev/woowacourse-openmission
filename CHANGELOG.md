#### 1.0.0 (2025-11-23)

##### Documentation Changes

* **readme:**
  *  문제 요구사항 수정 ([78290344](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/78290344b5aa05d7669b230012cd6244622e7f0a))
  *  미션 내용 작성 ([7c10888f](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/7c10888fdc34e3a2631da0a17edb89e7ead6bbe6))
*  도커 파일 추가 ([cb3a0914](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/cb3a091475d42d9975712a3e8e76bff97e92c31e))

##### New Features

*  요청 실패시 예외처리 핸들러 구현 ([5442f272](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/5442f27263933973e845c1df3b99e4096a7460b6))
*  상품 그룹화 가격 소수점 아래 3자리 수에서 반올림 ([1428c5eb](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/1428c5eb1e541f0de27fc6528f17efdd08133a0e))
*  상품 그룹들 한번에 결제하는 api구현 ([09d0d6e2](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/09d0d6e24f712897651913a1fa1f19537b33e3ac))
*  비동기 결제 api구현, 연결 ([8c1ea84f](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/8c1ea84f4264926930cd7e1c38682b1f6d75f9b0))
*  달러 환율 조회 api 연결 ([db7877cc](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/db7877cc81adc9f79e7415a00d162e4eb052645d))
*  상품 분할후 할인적용 api 연결 ([5ce52902](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/5ce5290255200dc2a5d0cc06eef8532a01ff8b83))
*  할인정보 조회 api 연결 ([c3387460](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/c3387460518e45cf38c0a25c996d184a74676dc6))
*  할인정보 조회 api구현 ([8562d365](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/8562d365f260dedad8a09e2c1e367f353b39cd00))
*  상품개수와 상품정보 입력 구현 ([94094b94](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/94094b94eb48da10e2a9bb28941005eae080d2a0))
*  환율 정보 조회 api 구현 ([2ee50fb9](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/2ee50fb9e72e2e326f15cd4e33830b2329d75d21))
*  최저가로 나누는 그룹핑 api구현 ([b4264afc](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/b4264afcc023acc73e66821a137c9852cce29d93))
*  상품 그룹화 알고리즘 구현, 할인정보 go코드로 변경 ([27c3bb01](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/27c3bb016cd0166deddb45e7b0428f4fe6d3ef10))
*  go 서버 프로잭트 생성 ([1782c60c](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/1782c60cc643fbb836a623bd94e7fb90837a224d))
* **policy:**  할인 정책 구현 ([0ae92b3f](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/0ae92b3fda23a06c3b454e3d081330e7af8346e1))

##### Bug Fixes

*  결제 시간 랜덤값으로 변경 ([52b5b343](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/52b5b34399349898f659eaa93d4b6324497f673b))
*  환율을 가져올수 있는 가장 최근날짜의 환율을 가져오는걸로 수정 ([fdac6fc4](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/fdac6fc437585242e23909d1a3089b15a99e5585))

##### Other Changes

* //github.com/Fiddich-Dev/woowacourse-openmission ([d1a7c5cf](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/d1a7c5cf47bbd3b5337e64f562176ce41caac3d0))

##### Refactors

*  PurchaseController로 이름 변경 ([28aecc15](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/28aecc150f90ec5670006aeae2f88ef6ee2f76da))
*  상수 분리 ([45649ef3](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/45649ef3591b48b6fbe10189f3de6d92bce91b71))
*  상수 처리 ([5e368bfb](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/5e368bfb8bd57914f507e13e0e734dd5266eacc3))
*  handler, service 로직들 코드 최적화 ([d767176c](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/d767176c6a1f2e1483008e1e1aa9f464a90a117d))
*  api 응답 포멧 수정 ([e4b1637b](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/e4b1637bc4958bc2abdd62b1003da4a5f7f57f54))
*  gin에 맞는 폴더구조와 네이밍으로 변경, api주소 변경 ([bf82c713](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/bf82c71366a8ac53b1ff34a637d645d581385a30))
*  쿠폰, 카드 할인 정책을 인베딩을 사용해 분리 ([18934cb3](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/18934cb3916ae9193d84b6b5129ef10ea0963f18))
*  Goods를 담는 컬렉션 이름을 goodsList로 변경 ([5c3aae98](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/5c3aae987c46f83c63637f5148f6854dd5ab0ce4))
*  api 요청 코드 모두 정적메서드로 변경, json 변환코드와 http 요청코드 분리 ([22091501](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/22091501dcc48d6034311e60c49dfce1e7c9331a))
*  할인정책에서 할인하는 메서드 제거 ([8c8cd2da](https://github.com/Fiddich-Dev/woowacourse-openmission/commit/8c8cd2dad525ce91e3dc3de1eb7a605bdf550d8d))


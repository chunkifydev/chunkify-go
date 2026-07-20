# Changelog

## 0.15.0 (2026-07-18)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/chunkifydev/chunkify-go/compare/v0.14.0...v0.15.0)

### Features

* **stlc:** configurable CI runner and private-production-repo support in workflow templates ([cd9b8aa](https://github.com/chunkifydev/chunkify-go/commit/cd9b8aacafad2c2e8bd8229703aacfc1c8724e39))

## 0.14.0 (2026-07-03)

Full Changelog: [v0.13.3...v0.14.0](https://github.com/chunkifydev/chunkify-go/compare/v0.13.3...v0.14.0)

### Features

* **client:** optimize json encoder for internal types ([258c291](https://github.com/chunkifydev/chunkify-go/commit/258c291130327662d38710c7495f0fc9a53cc8df))
* **go:** add default http client with timeout ([c243f27](https://github.com/chunkifydev/chunkify-go/commit/c243f273a978b25b53c50734a7aac63337569131))
* **internal:** support comma format in multipart form encoding ([8fac06a](https://github.com/chunkifydev/chunkify-go/commit/8fac06a5b34c4d9e857e834d0cf2ee7d304b2c4a))
* support setting headers via env ([286a9c6](https://github.com/chunkifydev/chunkify-go/commit/286a9c6e10dbc61647d0b885e1cf0b4be41f9034))


### Bug Fixes

* better respect format tags from the spec ([b3b5476](https://github.com/chunkifydev/chunkify-go/commit/b3b5476866a21bcd0561e2ae958ae289f9edbfc3))
* **go:** avoid panic when http.DefaultTransport is wrapped ([ded89a3](https://github.com/chunkifydev/chunkify-go/commit/ded89a3ef0a662dc8d3090cbff4a2e05714f721d))
* prevent duplicate ? in query params ([bded98f](https://github.com/chunkifydev/chunkify-go/commit/bded98f59de3247ae9797926bb28a2f9e56b74a6))


### Chores

* avoid embedding reflect.Type for dead code elimination ([69e7da2](https://github.com/chunkifydev/chunkify-go/commit/69e7da217b764778e2e8b3a213881c5203339c92))
* **ci:** skip lint on metadata-only changes ([94c3a62](https://github.com/chunkifydev/chunkify-go/commit/94c3a62f0aaa590b55c620633b39d1ed3a047e3b))
* **ci:** support opting out of skipping builds on metadata-only commits ([a2422af](https://github.com/chunkifydev/chunkify-go/commit/a2422af375258cc64b728897f30ab6195770bf93))
* **client:** fix multipart serialisation of Default() fields ([dd4fe59](https://github.com/chunkifydev/chunkify-go/commit/dd4fe599d22e9187301b51090395b0d3da12240d))
* **internal:** codegen related update ([e12eb6d](https://github.com/chunkifydev/chunkify-go/commit/e12eb6d71207d58e816313e05deec1abde67914c))
* **internal:** codegen related update ([b975581](https://github.com/chunkifydev/chunkify-go/commit/b9755813e8aa321a797abcfbc13a785e8f5cd7b4))
* **internal:** codegen related update ([bbdbad8](https://github.com/chunkifydev/chunkify-go/commit/bbdbad8a5bccbdf5e03cb465effe9a176d23c932))
* **internal:** codegen related update ([5352666](https://github.com/chunkifydev/chunkify-go/commit/5352666f4f88933d075d67e4344dbd794eff8bf9))
* **internal:** more robust bootstrap script ([b4ea570](https://github.com/chunkifydev/chunkify-go/commit/b4ea5701628f1d3369028b0dade5984074e7f60c))
* **internal:** support default value struct tag ([0f533ff](https://github.com/chunkifydev/chunkify-go/commit/0f533ff3ba0ca7c8f41ebe5597d680c6b765fc75))
* **internal:** tweak CI branches ([68ada7d](https://github.com/chunkifydev/chunkify-go/commit/68ada7d474fd05eb346f6249f5adfd46f07b6010))
* **internal:** update gitignore ([dd39b9e](https://github.com/chunkifydev/chunkify-go/commit/dd39b9e825354e07a7b04f4f3ed1040986dcb0da))
* redact api-key headers in debug logs ([fc37d82](https://github.com/chunkifydev/chunkify-go/commit/fc37d8286536fe51d2b7d412f09b739da626ab80))
* remove unnecessary error check for url parsing ([4bb109f](https://github.com/chunkifydev/chunkify-go/commit/4bb109f8423d273a736ebaedc1295c569ec90858))
* update docs for api:"required" ([d8ef05e](https://github.com/chunkifydev/chunkify-go/commit/d8ef05e4d214330a5335f6165b0140de434a5fd1))

## 0.13.3 (2026-03-11)

Full Changelog: [v0.13.2...v0.13.3](https://github.com/chunkifydev/chunkify-go/compare/v0.13.2...v0.13.3)

### Chores

* **internal:** codegen related update ([542479c](https://github.com/chunkifydev/chunkify-go/commit/542479c8f63e5cf767347c42c14167469be53a03))
* **internal:** minor cleanup ([d15a5bc](https://github.com/chunkifydev/chunkify-go/commit/d15a5bcc9e757df653a76bf9e1b28dbc96cb24dc))
* **internal:** use explicit returns ([199b41e](https://github.com/chunkifydev/chunkify-go/commit/199b41e21e9336ed69b4ea92ff7c18d7de079580))
* **internal:** use explicit returns in more places ([e10d8e6](https://github.com/chunkifydev/chunkify-go/commit/e10d8e6bb0988ab4867ade587c66ff59a8c88945))

## 0.13.2 (2026-03-07)

Full Changelog: [v0.13.1...v0.13.2](https://github.com/chunkifydev/chunkify-go/compare/v0.13.1...v0.13.2)

### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([ccacefd](https://github.com/chunkifydev/chunkify-go/commit/ccacefd40348a3c8ac3e59717336ff32eae2405d))
* **internal:** codegen related update ([38a77ef](https://github.com/chunkifydev/chunkify-go/commit/38a77efca721c315fdba620179bc6e806f992abe))

## 0.13.1 (2026-03-03)

Full Changelog: [v0.13.0...v0.13.1](https://github.com/chunkifydev/chunkify-go/compare/v0.13.0...v0.13.1)

### Chores

* **internal:** codegen related update ([38831c5](https://github.com/chunkifydev/chunkify-go/commit/38831c5b2c4973ff03aab3ade358bdbf9d8de079))
* **tests:** update webhook tests ([890485b](https://github.com/chunkifydev/chunkify-go/commit/890485bde76c883039b70fee7d3e712e10c0606d))

## 0.13.0 (2026-02-25)

Full Changelog: [v0.12.1...v0.13.0](https://github.com/chunkifydev/chunkify-go/compare/v0.12.1...v0.13.0)

### Features

* **api:** api update ([7c7529c](https://github.com/chunkifydev/chunkify-go/commit/7c7529c58b276597f7789d3b06e5b3e4ab7337e2))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([67a792e](https://github.com/chunkifydev/chunkify-go/commit/67a792ec0cf8fee9a1262c9655b9b1cf3ab02136))
* **encoder:** correctly serialize NullStruct ([9c568f1](https://github.com/chunkifydev/chunkify-go/commit/9c568f1e98d4bb576ecbdaa1156c12c6eb4aa7be))
* **internal:** skip tests that depend on mock server ([1bbe749](https://github.com/chunkifydev/chunkify-go/commit/1bbe7495afd77f83e4b7a1999b8499f2fd18ffd9))


### Chores

* **internal:** codegen related update ([d46dd67](https://github.com/chunkifydev/chunkify-go/commit/d46dd67058a10781590a741f163adad16a866156))
* **internal:** move custom custom `json` tags to `api` ([e184500](https://github.com/chunkifydev/chunkify-go/commit/e184500a32441b158f694b26c212d3998884cbaa))
* **internal:** remove mock server code ([cc0e255](https://github.com/chunkifydev/chunkify-go/commit/cc0e25579cb531d0085bc29ab200623283418248))
* update mock server docs ([eb35856](https://github.com/chunkifydev/chunkify-go/commit/eb35856e1c0dd3fdf1e3db616fdd5f349e91f11c))

## 0.12.1 (2026-02-05)

Full Changelog: [v0.12.0...v0.12.1](https://github.com/chunkifydev/chunkify-go/compare/v0.12.0...v0.12.1)

### Bug Fixes

* **client:** send correct authentication methods ([56ff2c4](https://github.com/chunkifydev/chunkify-go/commit/56ff2c4cc21832d66f6f1f4d7f89dc4f15732fac))


### Chores

* **internal:** codegen related update ([08e2c75](https://github.com/chunkifydev/chunkify-go/commit/08e2c75ccb9604a627682a0103a4d4db153edcca))

## 0.12.0 (2026-01-24)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/chunkifydev/chunkify-go/compare/v0.11.0...v0.12.0)

### Features

* **client:** add a convenient param.SetJSON helper ([a1ff01a](https://github.com/chunkifydev/chunkify-go/commit/a1ff01aa6961bf3015f86db425a090ae233d1fb0))

## 0.11.0 (2026-01-17)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/chunkifydev/chunkify-go/compare/v0.10.0...v0.11.0)

### Features

* **api:** api update ([f1834bc](https://github.com/chunkifydev/chunkify-go/commit/f1834bcd0c7994e51547f814b42eb7308858b00e))


### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([6403cd1](https://github.com/chunkifydev/chunkify-go/commit/6403cd147b832f3076964735253993b09cc8584b))


### Chores

* **internal:** update `actions/checkout` version ([e1fb1c4](https://github.com/chunkifydev/chunkify-go/commit/e1fb1c453f8990408b6e49598debf44187884d72))

## 0.10.0 (2026-01-15)

Full Changelog: [v0.9.3...v0.10.0](https://github.com/chunkifydev/chunkify-go/compare/v0.9.3...v0.10.0)

### Features

* **config:** added per endpoint security settings ([7701e0b](https://github.com/chunkifydev/chunkify-go/commit/7701e0b5f621f18992ffcddca6039372634f3701))

## 0.9.3 (2026-01-14)

Full Changelog: [v0.9.2...v0.9.3](https://github.com/chunkifydev/chunkify-go/compare/v0.9.2...v0.9.3)

### Chores

* **internal:** codegen related update ([44bf73b](https://github.com/chunkifydev/chunkify-go/commit/44bf73b65cf5a6b2a73d6e6d616f1a2592220a75))
* **sdk/config:** change model api_file to job-file ([2f2fce8](https://github.com/chunkifydev/chunkify-go/commit/2f2fce87bfa97c540dd66d9c4f4f4c64826df9d8))

## 0.9.2 (2026-01-06)

Full Changelog: [v0.9.1...v0.9.2](https://github.com/chunkifydev/chunkify-go/compare/v0.9.1...v0.9.2)

### Chores

* **internal:** codegen related update ([c3cc72f](https://github.com/chunkifydev/chunkify-go/commit/c3cc72f70a1b1954c79243e5632a30a58343cf3d))

## 0.9.1 (2025-12-19)

Full Changelog: [v0.9.0...v0.9.1](https://github.com/chunkifydev/chunkify-go/compare/v0.9.0...v0.9.1)

### Chores

* add float64 to valid types for RegisterFieldValidator ([ef0ee77](https://github.com/chunkifydev/chunkify-go/commit/ef0ee777920b179582316650af6328ce1f52cb92))
* **internal:** codegen related update ([bb0a087](https://github.com/chunkifydev/chunkify-go/commit/bb0a0875ef4e028ef508b95decb48935064b76dc))

## 0.9.0 (2025-12-18)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/chunkifydev/chunkify-go/compare/v0.8.0...v0.9.0)

### Features

* **api:** manual updates ([12251f5](https://github.com/chunkifydev/chunkify-go/commit/12251f5cb83339e8038c8b36a6b934c2afee35ae))


### Bug Fixes

* skip usage tests that don't work with Prism ([94f1ee5](https://github.com/chunkifydev/chunkify-go/commit/94f1ee5144b795d38c355d4b9b5bc254862ef350))

## 0.8.0 (2025-12-12)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/chunkifydev/chunkify-go/compare/v0.7.0...v0.8.0)

### Features

* **encoder:** support bracket encoding form-data object members ([e2e2496](https://github.com/chunkifydev/chunkify-go/commit/e2e2496e6fcf9514f3c52cd3e8ee2c42dad2cd47))


### Bug Fixes

* **mcp:** correct code tool API endpoint ([2e70115](https://github.com/chunkifydev/chunkify-go/commit/2e7011536e04735858e5569eac2d61bcfad75094))
* rename param to avoid collision ([e5e4d9c](https://github.com/chunkifydev/chunkify-go/commit/e5e4d9c2ed01610da12d4ad1f2027cd44f6a585f))


### Chores

* elide duplicate aliases ([7e9a6a2](https://github.com/chunkifydev/chunkify-go/commit/7e9a6a25a69d9238e828b19fdc7241515215cacc))
* **internal:** codegen related update ([d5d2b40](https://github.com/chunkifydev/chunkify-go/commit/d5d2b40b0f917d40807b4d43ca0d2be86ce1131f))
* **internal:** codegen related update ([4f45b16](https://github.com/chunkifydev/chunkify-go/commit/4f45b1690dd2b165ed2e868c4c242a694f903c1b))

## 0.7.0 (2025-12-01)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/chunkifydev/chunkify-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** manual updates ([cd47b10](https://github.com/chunkifydev/chunkify-go/commit/cd47b101a44a4cd2cbb54cd9f2876b67545e27cf))

## 0.6.0 (2025-11-27)

Full Changelog: [v0.3.0...v0.6.0](https://github.com/chunkifydev/chunkify-go/compare/v0.3.0...v0.6.0)

### Features

* **api:** manual updates ([d909424](https://github.com/chunkifydev/chunkify-go/commit/d909424baa6717640fcc4d673c2e6b5eab195126))

## 0.3.0 (2025-11-27)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/chunkifydev/chunkify-go/compare/v0.2.0...v0.3.0)

### ⚠ BREAKING CHANGES

* **api:** update all created.* query string to epoch unix time format

### Features

* **api:** update all created.* query string to epoch unix time format ([7aa5f40](https://github.com/chunkifydev/chunkify-go/commit/7aa5f408b15e374193a3318e68e82981117de840))

## 0.2.0 (2025-11-26)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/chunkifydev/chunkify-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([47011d3](https://github.com/chunkifydev/chunkify-go/commit/47011d39c7d30b1d8041a50b6a649d49f3cbaf4e))

## 0.1.0 (2025-11-26)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/chunkifydev/chunkify-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([245bf9a](https://github.com/chunkifydev/chunkify-go/commit/245bf9ae48d95e0a751347c83cb2813b20898a1b))
* **api:** manual updates ([f9b637f](https://github.com/chunkifydev/chunkify-go/commit/f9b637f646b75adf80490cba1611932049db031e))
* **api:** manual updates ([f94a60d](https://github.com/chunkifydev/chunkify-go/commit/f94a60d92a1c34956f0d063ddc39553239cb1ecf))
* **api:** manual updates ([3a24922](https://github.com/chunkifydev/chunkify-go/commit/3a249226f422b8ab19b6da34a58dc26d7f0593e9))
* **api:** manual updates ([6084fcb](https://github.com/chunkifydev/chunkify-go/commit/6084fcb25da926deea76ce0e126dbdc43231a86b))
* **api:** manual updates ([ab5184c](https://github.com/chunkifydev/chunkify-go/commit/ab5184ca1357d86f9e0dd1df7ee0543fac7e79cb))
* **api:** manual updates ([937ac24](https://github.com/chunkifydev/chunkify-go/commit/937ac2409ca0a5ec4b77220e620302265b480287))
* **api:** manual updates ([3a99684](https://github.com/chunkifydev/chunkify-go/commit/3a99684d9ce3853e3666a33c53db53d3350d951f))
* **api:** manual updates ([761dd60](https://github.com/chunkifydev/chunkify-go/commit/761dd6054ef1317c60a23553e1a08b55eca253f6))
* **api:** manual updates ([3968042](https://github.com/chunkifydev/chunkify-go/commit/3968042c927acb8a0c064e698907c7a7853b5660))
* **api:** manual updates ([ef12faf](https://github.com/chunkifydev/chunkify-go/commit/ef12faf866d2e246ca5dfdb8186a9e9aaa9fc306))
* **api:** manual updates ([53bddd4](https://github.com/chunkifydev/chunkify-go/commit/53bddd41296bc44a5905d9fbe51a80e76cf8831b))
* **api:** manual updates ([c5e8734](https://github.com/chunkifydev/chunkify-go/commit/c5e873414780715193ff8238780d07d72aa49a64))
* **api:** manual updates ([6dda26c](https://github.com/chunkifydev/chunkify-go/commit/6dda26c897c8a4d82488860900750eee9f32d46c))
* **api:** manual updates ([5e51dea](https://github.com/chunkifydev/chunkify-go/commit/5e51deaf3310cce2279655a670b17fea5a4f157a))
* **api:** manual updates ([4f85223](https://github.com/chunkifydev/chunkify-go/commit/4f85223cc683a6feb6744ee2184a9198181a2597))
* **api:** manual updates ([821819b](https://github.com/chunkifydev/chunkify-go/commit/821819b8401a3dc94fd734b7059d18568167f202))
* **api:** manual updates ([4c9911d](https://github.com/chunkifydev/chunkify-go/commit/4c9911da3c14b3078607bd5d5905ad2c55b269ac))
* **api:** manual updates ([ba1ada0](https://github.com/chunkifydev/chunkify-go/commit/ba1ada0a7bd100de9b5730693d9ed37fd371efd9))
* **api:** manual updates ([1d22f37](https://github.com/chunkifydev/chunkify-go/commit/1d22f370efbbe19b049c66c813f140504e0bc96e))
* **api:** manual updates ([4986501](https://github.com/chunkifydev/chunkify-go/commit/498650131e593e8b9bee0838e0b0aa5955250d0d))
* **api:** manual updates ([0bcd356](https://github.com/chunkifydev/chunkify-go/commit/0bcd356888ec5337b89503184dcf631bc6f3ce3b))
* **api:** manual updates ([2cbf318](https://github.com/chunkifydev/chunkify-go/commit/2cbf318b758b5aa78ab28e3ed37393e914f68191))
* **api:** manual updates ([3fed433](https://github.com/chunkifydev/chunkify-go/commit/3fed433cb6599c9cea5b2433f3fdc7a58da8c3b3))
* **api:** manual updates ([b82bf54](https://github.com/chunkifydev/chunkify-go/commit/b82bf54a2d33ab9f25f3d40ebc9dbc2fe6f92025))
* **api:** manual updates ([4c00857](https://github.com/chunkifydev/chunkify-go/commit/4c00857a500f64ca927e43e07cd2cc59c94bad38))
* **api:** manual updates ([0becd77](https://github.com/chunkifydev/chunkify-go/commit/0becd77fa943fb81d3ecce85bdfec2c30a1de36a))
* **api:** manual updates ([8e67d8f](https://github.com/chunkifydev/chunkify-go/commit/8e67d8f6c8e26e97cb414ae091109ef6cdd6fa14))
* **api:** manual updates ([57f9144](https://github.com/chunkifydev/chunkify-go/commit/57f9144179bb56844581f806380441b49b276146))
* **api:** manual updates ([f3929f4](https://github.com/chunkifydev/chunkify-go/commit/f3929f47d8aba628e83e039b943704549a9d5279))
* **api:** manual updates ([1cd6f82](https://github.com/chunkifydev/chunkify-go/commit/1cd6f82e3c54ce9739be81b8fd3216ad78dcdfb4))
* **api:** manual updates ([40e38dd](https://github.com/chunkifydev/chunkify-go/commit/40e38dd9103fb9e90db7071b051a7c4be187f474))
* **api:** manual updates ([cd20d6c](https://github.com/chunkifydev/chunkify-go/commit/cd20d6c18b6d3a2c75b5cdba92714e115f2ec1c1))
* **api:** manual updates ([cbba971](https://github.com/chunkifydev/chunkify-go/commit/cbba97102d2eed9d2d0b51a090a9e4bce391c339))
* **api:** manual updates ([585f711](https://github.com/chunkifydev/chunkify-go/commit/585f7118dcbadf016d9abaf665fab876228cca19))
* **api:** manual updates ([2c02182](https://github.com/chunkifydev/chunkify-go/commit/2c021829e8a88ebe6991abdb146c6062f8535bc6))
* **api:** manual updates ([6724159](https://github.com/chunkifydev/chunkify-go/commit/672415948aac7dee76dbad7b5dc228101ef86254))
* **api:** manual updates ([f42aabd](https://github.com/chunkifydev/chunkify-go/commit/f42aabd5c28b18c83fe7c2c4b2d0b00f979c9d18))
* **api:** manual updates ([192f87c](https://github.com/chunkifydev/chunkify-go/commit/192f87c04529b00449309afe0745a0975623504a))
* **api:** manual updates ([fa6ea9f](https://github.com/chunkifydev/chunkify-go/commit/fa6ea9f1bbaf9f668b1eca5c47a16ce6580fd3d9))
* **api:** manual updates ([efea48f](https://github.com/chunkifydev/chunkify-go/commit/efea48fc69cb41df761605eeb44052af52a706ba))
* **api:** manual updates ([46fd8d5](https://github.com/chunkifydev/chunkify-go/commit/46fd8d560a02e64325c34b9265d9fb47916721aa))
* **api:** manual updates ([c425675](https://github.com/chunkifydev/chunkify-go/commit/c425675a404ca80110bac4a18411e9a08746d91d))
* **api:** manual updates ([7c221dd](https://github.com/chunkifydev/chunkify-go/commit/7c221dd58c34ccf85efde75392ef89af3173f0c5))
* **api:** manual updates ([be38d96](https://github.com/chunkifydev/chunkify-go/commit/be38d96df6cb1912ea041b7d480b3359cf265358))
* **api:** manual updates ([bb4315a](https://github.com/chunkifydev/chunkify-go/commit/bb4315a4d1cf7f081cffae587795e56eae47339f))
* **api:** manual updates ([3fd0430](https://github.com/chunkifydev/chunkify-go/commit/3fd04308c117cd996271ce12031e2aca328a5ed4))
* **api:** manual updates ([bd93812](https://github.com/chunkifydev/chunkify-go/commit/bd93812fb7b53b0d9a93f5f0bdda2d07239fb672))
* **api:** manual updates ([38efd66](https://github.com/chunkifydev/chunkify-go/commit/38efd66d4b583639d7d27acd16e510a03e340438))
* **api:** manual updates ([3ce9c3c](https://github.com/chunkifydev/chunkify-go/commit/3ce9c3ca4a22fc061dd6fed9deb5fdd061206098))
* **api:** manual updates ([726210f](https://github.com/chunkifydev/chunkify-go/commit/726210f860ad147e354ac8810dde29420f738cbb))
* **api:** manual updates ([c1f792e](https://github.com/chunkifydev/chunkify-go/commit/c1f792e6a511b73adc428e89d80da52313e89330))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([e900eee](https://github.com/chunkifydev/chunkify-go/commit/e900eee7e585fe1c89f214c9ba921a4688ce55dc))


### Chores

* update SDK settings ([b69dfff](https://github.com/chunkifydev/chunkify-go/commit/b69dfffa2c4da534fa823aedfa85c0802835c292))

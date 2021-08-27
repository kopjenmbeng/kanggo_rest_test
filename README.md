# kanggo_rest_test

Services ini dibuat untuk soal berikut ![klik disini untuk melihat soal](../master/document/SoaltestITBackEnd)

# Technology Stack
-   GO
-   Postgres
-   Consul untuk store configuration (Optional)
-   Postman untuk API Documentation

# Sebelum RUN
Silahkan restore database berikut ![klik untuk download](../master/script/kanggo_db.sql)

-   pastikan .env file ini ada and the konfigurasinya seperti ini.
    # predefined goconf env vars
    - GOCONF_ENV_PREFIX=kanggo
    - #GOCONF_CONSUL=localhost:8500 (please remark using # if you don't have consul so it will read ![this json config](../master/kanggo.config.json))
    - GOCONF_TYPE=json
    - GOCONF_FILENAME=kanggo.config

    # Newrelic
    - #PROPERTY_NEWRELIC_KEY=

setelah memenuhi permintaan diatas ,berikut cara menjalankan service ini secara local .
-   untuk menjalankan api silahkan jalankan command berikut
    -   go run main.go api  maka aplikasi akan jalan di http://localhost:8080
-   for Api documentation silahkan import ![api doc](../master/document/kanggo.postman_collection.json))

-   untuk mencoba api ini secara publik silahkan gunakan base url berikut.
    -   kanggo-api.slametsupriyadi.com


Feel free for ask
-   WA 087777000056
-   email slamet.supriyadi88@gmail.com

Regard
## Noted : 
Saya tidak menggunakan cluster k8s pribadi sehingga tidak ada proses deployment ke cluster, saya hanya integrasi dari github ke cluter dan melakukan test koneksi. jadi saya mohon untuk tidak mengotak-atik repo dan share repo ke orang lain karena ini merupakan tanggup jawab saya. Terimakasih

## Getting started
###### Cicd berjalan di branch staging
1. Membuat simple App menggunakan Go di file main.go, app ini hanya menampilkan Text
2. Pada awalnya saya ingin menerapkan cicd menggunakan tools jenkins dengan konsep shared library. Namun karena keterbatasan server jenkins jadi saya menggunakan github action sebagai pengganti.
3. Untuk menggunakan github action diperlukan direktory .github/workflows/ sebagai tempat menyimpan scirpt pipeline. Dalam hal ini saya membuat beberapa stage seperti checkoutcode, build images, push images, deploy to k8s. Semua stage saya simpan pada file cicd.yml di direktory .github/workflows/.
4. Sebelum menjalankan stage build dan push images, saya membuat Dockerfile yang berisi perintah untuk menginstall dan menjalankan environment yang dibutuhkan untuk menjalankan simple Go app. yang kemudian akan di-build mengjadi images dan di push ke images registry di dockerhub
5. Kemudian untuk dapat login ke dockerhub registry dibutuh access berupa credential username/password. Dalam hal ini saya setuo secret untuk username dan password saya yang kemudian akan saya deklarasikan di stage login to dockerhub
6. kemudian di stage build dan push images, Dockerfile akan di build terlebih dulu kemudian di push ke docker hub sebagai images registry. ditahap ini saya melakukan penaman latest dengan taging sesuai dengan build number.
7. Kemudian ke stage Configure ke kubernetes, dimana saya meng-integrasikan pipeline ke kubernetes. Ditahap ini saya memerlukan secred-config berupa base64. Saya meng-integrasikannya dengan k8s yang terinstall di Digital Oceon
8. Dan untuk stage deploy to k8s, saya hanya melakukan check namespace untuk confirm bahwa sya sudah bisa koneksi ke cluster. untuk deploy manifestnya, saya tidak lanjutkan karena cluster yang saya gunakan bukanlah milik saya tapi meminjam dari teman sehingga akses terbatas. namun sebagai rujukan saya membuat contoh proses deployment manifest di pipeline yang sudah sya beri argumen(comment). tahap deployment manifestnya saya buat dalam dua cara yaitu manual menggunakan file deployment.yaml/service.yaml di direktory manifest dan menggunakan templete helm-chart di direktory sample.
9. saya juga tidak apply atau setup domain karena app saya tidak berjalan di cluster, tapi untuk domain biasanya saya menggukan ingressClass yang di setup di helm chart atau secara manual bisa menambahkan file ingress.yaml. Saya juga pernah menggunakan layanan DoDaddy sebagai domain name

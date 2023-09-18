#Run build image <br />
cd to _infra <br />
docker-compose build <br />
docker-compose up <br />

start dynamodb-local first
start my_service-1 second

#Default on http://localhost:8081

#Testing Example
1. call post "/userprofile/save" for create user profile <br />
{
    "userName": "d",
    "name":"robin",
    "createdBy" : "admin",
    "email": "d@d.com"
}

2. keep userReference from 1. <br />

3. call post "/appointment/save" for save appointment card โดย เอา userReference จาก 1. ที่ return มาใช้ <br />
 ถ้า update appointment ให้ส่ง appId ของ appointment นั้น ไปด้วย <br />
#<b>example create</b> <br />
{
    "title": "นัดหมาย2",
    "description": "ทดสอบ 23",
    "createdBy": "den",
    "email": "d@d.com",
    "userReference": "{userReference}",  
    "status" : "todo"
}<br />
#<b>example update</b> <br />
{
    "appId": {appId},
    "title": "นัดหมาย2",
    "description": "ทดสอบ 456",
    "updatedBy": "den",
    "email": "d@d.com",
    "userReference": "{userReference}",  
    "status" : "in_progress"
}

4. call get appointment list from "/appointment/list?lastKey={appId}&limit=10" <br />
    lastKey = appId key จากข้อมูลตัวล่าสุดที่จะเข้าไปโหลดข้อมูลเพิ่ม ไม่ใส่ จะ เอาตัวแรกสุด <br />
    limit = จำนวน record ที่จะนำมาแสดง ถ้าไม่ใส่ จะ default = 10 <br />

5. call post "/appointment/comment/save" for save comment appId ดูตาม appointment card ที่ต้องการเพิ่ม comment <br />
ถ้า update comment ให้ส่ง id ของ comment นั้น ไปด้วย <br />
<b>#example create</b> <br />
{
    "appId": {appId},
    "message": "ทดสอบ comment",
    "createdBy": "robinhood",
    "userReference" : {userReference}
}<br />
<b>#example update</b> <br />
{
    "id" : {id},
    "appId": {appId},
    "message": "ทดสอบ comment",
    "updatedBy": "robinhood",
    "userReference" : {userReference}
}

6. call get appointment detail (ดู comments ทั้งหมด) by appointment card from "/appointment/detail?appId={appId}" <br />
 appId ดูตาม appointment card จาก 4. ที่ต้องการ ดูเนื้อหาข้างใน 

7. get master status from /master-data/status

#for unit test
run docker run -p 4566:4566 localstack/localstack

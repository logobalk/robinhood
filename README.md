#Run build image
cd to _infra
docker-compose build
docker-compose up

#Testing Example
1. call post /userprofile/save for create user profile
{
    "userName": "d",
    "name":"robin",
    "createdBy" : "admin",
    "email": "d@d.com"
}

2. keep userReference from 1.
3. call post /appointment/save for save appointment card ถ้า update ให้ส่ง appId ไปด้้วย
{
    "title": "นัดหมาย2",
    "description": "ทดสอบ 23",
    "createdBy": "den",
    "email": "d@d.com",
    "userReference": "7d307afd-e8c9-486f-8098-a231bbe568d0",  // user userReference from 2.
    "status" : "todo"
}

4. check appointment from /appointment/list?lastKey=22cb93c2-a6ef-4834-b003-a0620c3968dd&limit=10
    lastKey = key จากข้อมูลตัวล่าสุดที่จะเข้าไปโหลดข้อมูลเพิ่ม ไม่ใส่ จะ เอาตัวแรกสุด
    limit = จำนวน record ที่จะนำมาแสดง ถ้าไม่ใส่ จะ default = 10

5. call post /appointment/comment/save for save comment ถ้า update ให้ส่ง id ไปด้้วย
{
    "appId": "985b7fa4-3b8a-40b4-b599-dc0750cf866c", // get from appointment/list
    "message": "ทดสอบ comment",
    "createdBy": "robinhood",
    "userReference" : "16d38871-9315-453a-9d76-c5a61981a966"
}

6. get master status from /master-data/status
# Uy vazifasi: Task Scheduler

## Topshiriq Tavsifi:
Foydalanuvchilarga ma'lum vaqtda bajarilishi kerak bo'lgan vazifalarni yaratish imkoniyatini beruvchi tizimni ishlab chiqing. Ushbu tizim takroriy vazifalarni avtomatlashtirish, eslatmalar yuborish yoki fon ishlarini bajarish uchun foydali bo'lishi mumkin.


## Talablar:
1. **Vazifa yaratish:**
    - Foydalanuvchilar vazifa tafsilotlarini va bajarilishi kerak bo'lgan vaqtni belgilash orqali vazifalar yaratishi mumkin.
    - Har bir vazifa quyidagilarni o'z ichiga oladi:
        - `task_name`: Vazifaning nomi.
        - `schedule_time`: Vazifa bajarilishi kerak bo'lgan vaqt.
        - `command`: Belgilangan vaqtda bajariladigan harakat yoki buyruq.

2. **Vazifalarni bajarish:**
    - Tizim vazifalarni belgilangan vaqtda bajarishi kerak
    - Takrorlanuvchi vazifalarni qo'llab-quvvatlash (masalan, har kuni soat 10:00 da).
    - Agar kerak bo'lsa, turli vaqt zonalarini hisobga olish imkoniyati.
    

3. **Vazifalarni boshqarish:**
    - Foydalanuvchilar rejalashtirilgan vazifalarni ko'rishi, tahrirlashi yoki o'chirishi mumkin.
    - Vazifaning holati (masalan, kutilmoqda, bajarilgan, xato) kuzatib boriladi.

4. **Bildirishnomalar:**
    - Vazifalar bajarilganda yoki xatolik yuz berganda bildirishnoma (elektron pochta yoki xabar) yuborish opsiyasi.

5. **Loglash:**
    - Har bir vazifa bajarilishining natijasi, jumladan, muvaffaqiyatli yoki muvaffaqiyatsizligi haqidagi ma'lumotni tizimda yozib borish.

6. **Qo'shimcha imkoniyatlar**:
    - `Cron`-ga o'xshash rejalashtirish: Foydalanuvchilar ko'proq murakkab rejalashtirish imkoniyatlari uchun `cron` ifodalarini belgilashlari mumkin.
    - Vazifalarni ustuvorlik bilan bajarish: Foydalanuvchilar vazifalarga ustuvorliklar berishi mumkin (masalan, yuqori, o‘rta, past), shunda tizim bir vaqtda rejalashtirilgan vazifalar orasida birinchi navbatda yuqori ustuvorlikdagi vazifalarni bajaradi.
    - Concurrency: Bir nechta vazifalar bir vaqtning o'zida bajarilishi mumkin bo'lishi va tizimni bloklamaslikka e'tibor berish.

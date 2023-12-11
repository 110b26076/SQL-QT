# 1.蒐集系統需求
**Entity**:配對者、被配對者  
**Relationship**:配對者與被配對者間有配對relationship  
**Attribute**:配對者{id、gender、height、age、ask_gender_up、ask_gender_low、ask_height_up、ask_height_low、ask_age_up、ask_age_low}、被配對者{id、gender、height、age、ask_gender_up、ask_gender_low、ask_height_up、ask_height_low、ask_age_up、ask_age_low}  
**Key**:配對者的id和被配對者的id  
**Cardinality**:配對者可配對0到多的被配對者，被配對者可被配對0到多的配對者  

# 2.ER Diagram
https://github.com/chu1111/database_project/tree/main/ER-Diagram  
![image](https://github.com/chu1111/database_project/blob/main/ER-Diagram/database.drawio.png)  

# 3.邏輯資料模型
https://github.com/chu1111/database_project/tree/main/Table  
![image](https://github.com/chu1111/database_project/blob/main/Table/tabe.drawio.png)

# 4.實體資料模型
```sql
CREATE TABLE matched_table (
    id INT PRIMARY KEY,
    gender VARCHAR(255),
    height INT,
    age INT,
    ask_gender VARCHAR(255),
    ask_height_up INT,
    ask_height_low INT,
    ask_age_up INT,
    ask_age_low INT
);
```
```Mysql
DELIMITER $$

CREATE PROCEDURE insert_data()
BEGIN
    DECLARE count INT DEFAULT 1;
    DECLARE GEN FLOAT;
	DECLARE ASK_HEIGHT FLOAT;
    DECLARE ASK_AGE float;
    
    WHILE count <= 100 DO
        SET GEN = RAND();
		SET ASK_HEIGHT = RAND();
        SET ASK_AGE  =  RAND();
        IF GEN > 0.5 THEN
            INSERT INTO matched_table(id, gender, height, age, ask_gender, ask_height_up, ask_height_low, ask_age_up, ask_age_low)
            VALUES (
                count, 'MAN', ROUND(RAND() * 10) + 160 + ROUND(GEN),
                ROUND(RAND() * 10) + 15, IF(RAND() > 0.2, 'woman', 'man'),
                ROUND(ASK_HEIGHT * 10) + 150 + ROUND(GEN * 1.5), ROUND(ASK_HEIGHT * 10) + 150,
                ROUND(ASK_AGE * 12) + 18 + 1, ROUND(ASK_AGE * 8) + 18
            );
        ELSE
            INSERT INTO matched_table(id, gender, height, age, ask_gender, ask_height_up, ask_height_low, ask_age_up, ask_age_low)
            VALUES (
                count, 'WOMAN', ROUND(RAND() * 10) + 150 + ROUND(GEN),
                ROUND(RAND() * 10) + 15, IF(RAND() > 0.2, 'man', 'woman'),
                ROUND(ASK_HEIGHT * 10) + 170 + round(GEN)  , ROUND(ASK_HEIGHT * 10) + 160,
                ROUND(ASK_AGE * 15) + 15, ROUND(ASK_AGE* 8) + 15
            );
        END IF;

        SET count = count + 1;
    END WHILE;
END $$

DELIMITER ;

CALL insert_data();
```

# 系統展示


https://github.com/chu1111/database_project/assets/113088657/49e1b3f8-0e5f-4047-ad1f-879da252c49d





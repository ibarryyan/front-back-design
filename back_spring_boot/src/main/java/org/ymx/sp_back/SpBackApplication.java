package org.ymx.sp_back;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("org.ymx.sp_back.xmlMapper") //扫描的mapper
public class SpBackApplication {

    public static void main(String[] args) {
        SpringApplication.run(SpBackApplication.class, args);
    }

}

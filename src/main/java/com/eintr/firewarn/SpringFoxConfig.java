package com.eintr.firewarn;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import springfox.documentation.builders.ApiInfoBuilder;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.oas.annotations.EnableOpenApi;
import springfox.documentation.service.ApiInfo;
import springfox.documentation.service.Contact;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;


/**
 * Swagger配置类
 */
@Configuration
@EnableOpenApi
public class SpringFoxConfig {
	@Bean
	public Docket docket(){
		return new Docket(DocumentationType.SWAGGER_2)
						.apiInfo(apiInfo()).enable(true)
						.select()
						//apis： 添加swagger接口提取范围
						.apis(RequestHandlerSelectors.basePackage("com.eintr"))
						//.apis(RequestHandlerSelectors.withMethodAnnotation(ApiOperation.class))
						.paths(PathSelectors.any())
						.build();
	}

	private ApiInfo apiInfo(){
		return new ApiInfoBuilder()
						.title("XX项目接口文档")
						.description("XX项目描述")
						.contact(new Contact("impact-eintr",
										"https://github.com/impact-eintr",
										"yixingwei4@gmail.com"))
						.version("1.0")
						.build();
	}
}
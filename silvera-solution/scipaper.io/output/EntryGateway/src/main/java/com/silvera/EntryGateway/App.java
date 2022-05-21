package com.silvera.EntryGateway;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;

import org.springframework.cloud.client.discovery.EnableDiscoveryClient;



@SpringBootApplication

@EnableDiscoveryClient
public class App {

  public static void main(String[] args) {
    SpringApplication.run(App.class, args);
  }

  @Bean
  public RouteLocator gatewayRoutes(RouteLocatorBuilder routeLocatorBuilder)
  {
    return routeLocatorBuilder.routes()
            .route("User", rt -> rt.path("/users/**")
                    .filters(f -> f.rewritePath("/users/(?<segment>.*)", "/${segment}"))
                    .uri("lb://User"))
    
            .route("SciPaper", rt -> rt.path("/papers/**")
                    .filters(f -> f.rewritePath("/papers/(?<segment>.*)", "/${segment}"))
                    .uri("lb://SciPaper"))
    
            .route("Library", rt -> rt.path("/library/**")
                    .filters(f -> f.rewritePath("/library/(?<segment>.*)", "/${segment}"))
                    .uri("lb://Library"))
    
            .build();
  }

}
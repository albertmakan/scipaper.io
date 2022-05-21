/**
    THIS IS GENERATED CODE AND SHOULD NOT BE CHANGED MANUALLY!!!

    Generated by: silvera
    Date: 2022-05-19 22:55:47
*/

package com.silvera.SciPaper.domain.model;

import org.springframework.data.annotation.Id;
import javax.validation.constraints.*;
public class Section {

    
    @Id
    private String id;
    
    
    
    @NotBlank(message="Field 'name' cannot be empty!")
    private java.lang.String name;
    
    
    @NotBlank(message="Field 'content' cannot be empty!")
    private java.lang.String content;


    
    public java.lang.String getName() {
        return this.name;
    }

    public void setName(java.lang.String name) {
        this.name = name;
    }
    
    public java.lang.String getContent() {
        return this.content;
    }

    public void setContent(java.lang.String content) {
        this.content = content;
    }
    

    
    public String getId(){
        return this.id;
    }
    

}
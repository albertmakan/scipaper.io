/**
    THIS IS GENERATED CODE AND SHOULD NOT BE CHANGED MANUALLY!!!

    Generated by: silvera
    Date: 2022-05-19 22:55:47
*/

package com.silvera.Library.messages;

import java.util.ArrayList;
import java.util.List;

public class Message {

    protected String name;
    protected List<MessageField> fields;
    protected List<MessageAnnotation> annotations;

    public Message(){}

    public Message(String name){
        this.name = name;
        this.fields = new ArrayList<>();
        this.annotations = new ArrayList<>();
    }

    public Message(String name, List<MessageField> fields, List<MessageAnnotation> annotations) {
        this.name = name;
        this.fields = fields;
        this.annotations = annotations;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public List<MessageField> getFields() {
        return fields;
    }

    public void setFields(List<MessageField> fields) {
        this.fields = fields;
    }

    public List<MessageAnnotation> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(List<MessageAnnotation> annotations) {
        this.annotations = annotations;
    }
}
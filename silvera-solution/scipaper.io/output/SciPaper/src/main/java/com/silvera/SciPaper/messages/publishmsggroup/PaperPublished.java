/**
    THIS IS GENERATED CODE AND SHOULD NOT BE CHANGED MANUALLY!!!

    Generated by: silvera
    Date: 2022-05-19 22:55:47
*/

package com.silvera.SciPaper.messages.publishmsggroup;

import com.silvera.SciPaper.messages.Message;

public class PaperPublished extends Message {

    
    private java.lang.String paperId;
    private java.lang.String author;
    private java.lang.String title;


    public PaperPublished(){
        this.name = "PublishMsgGroup.PaperPublished";
    }
    public java.lang.String getPaperId() {
        return this.paperId;
    }

    public void setPaperId(java.lang.String paperId) {
        this.paperId = paperId;
    }
    
    public java.lang.String getAuthor() {
        return this.author;
    }

    public void setAuthor(java.lang.String author) {
        this.author = author;
    }
    
    public java.lang.String getTitle() {
        return this.title;
    }

    public void setTitle(java.lang.String title) {
        this.title = title;
    }
    

}
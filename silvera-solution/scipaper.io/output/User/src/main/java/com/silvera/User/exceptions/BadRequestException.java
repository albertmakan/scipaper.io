package com.silvera.User.exceptions;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.BAD_GATEWAY)
public class BadRequestException extends RuntimeException {
  public BadRequestException() {
    super();
  }
  public BadRequestException(String message) {
      super(message);
  }
}

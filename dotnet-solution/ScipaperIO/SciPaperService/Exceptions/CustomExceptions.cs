using System;

namespace SciPaperService.Exceptions
{
    public class BadLogicException : Exception
    {
        public BadLogicException() : base() { }
        public BadLogicException(string message) : base(message) { }
    }


    public class NotFoundException : Exception
    {
        public NotFoundException() : base() { }
        public NotFoundException(string message) : base(message) { }
    }


    public class ForbiddenException : Exception
    {
        public ForbiddenException() : base() { }
        public ForbiddenException(string message) : base(message) { }
    }


    public class UnauthorizedException : Exception
    {
        public UnauthorizedException() : base() { }
        public UnauthorizedException(string message) : base(message) { }
    }

}

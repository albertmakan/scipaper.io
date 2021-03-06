using Microsoft.AspNetCore.Http;
using System;
using System.Net;
using System.Text.Json;
using System.Threading.Tasks;
using UserService.Exceptions;

namespace UserService.Middleware
{
    public class ErrorHandlerMiddleware
    {
        private readonly RequestDelegate _next;
        public ErrorHandlerMiddleware(RequestDelegate next)
        {
            _next = next;
        }

        public async Task InvokeAsync(HttpContext context)
        {
            try
            {
                await _next(context);
            }
            catch (Exception exception)
            {
                await HandleResponseBasedOnException(context, exception);
            }
        }

        private static Task HandleResponseBasedOnException(HttpContext context, Exception exception)
        {
            context.Response.ContentType = "application/json";
            context.Response.StatusCode = exception switch
            {
                NotFoundException => (int)HttpStatusCode.NotFound,
                BadLogicException => (int)HttpStatusCode.BadRequest,
                ForbiddenException => (int)HttpStatusCode.Forbidden,
                UnauthorizedException => (int)HttpStatusCode.Unauthorized,
                _ => (int)HttpStatusCode.InternalServerError,
            };
            var responseModel = new BaseApiResponse(context.Response.StatusCode, exception.Message);
            var result = JsonSerializer.Serialize(responseModel);
            return context.Response.WriteAsync(result);
        }
    }

    internal record struct BaseApiResponse(int StatusCode, string Message);
}

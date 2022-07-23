﻿using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;

namespace UserService.Utils
{
    public static class UrlHelpers
    {
        private static bool IsSimple(Type t)
            => t.IsPrimitive || t == typeof(string) || t == typeof(DateTime);
        private static string UriData(string key, object value)
            => $"{Uri.EscapeDataString(key)}={Uri.EscapeDataString(value.ToString())}";

        public static string ToQueryString(this object request, string name, int lvl = 0)
        {
            if (request == null || name == null) return string.Empty;

            if (request is DateTime date) return UriData(name, date.ToString("yyyy-MM-dd"));
            if (IsSimple(request.GetType())) return UriData(name, request);

            var items = new List<string>();

            if (request is IDictionary dict)
            {
                foreach (var key in dict.Keys)
                    if (dict[key] != null && IsSimple(dict[key].GetType()))
                        items.Add(UriData(key.ToString(), dict[key]));
            }
            else if (request is IEnumerable enumerable)
            {
                foreach (var item in enumerable)
                    if (item != null && IsSimple(item.GetType()))
                        items.Add(UriData(name, item));
            }
            else
                items = request.GetType().GetProperties()
                    .Where(x => x.CanRead && x.GetValue(request) != null)
                    .Select(x => ToQueryString(x.GetValue(request), $"{(name == "" || lvl < 1 ? "" : name + ".")}{x.Name}", lvl + 1))
                    .Where(qs => qs != string.Empty)
                    .ToList();

            return string.Join('&', items);
        }
    }
}

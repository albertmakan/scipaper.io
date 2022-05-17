import React from "react";
import * as yup from "yup";
import { toast } from "react-toastify";
import { login } from "../Services/userService";
import { TextField } from "@mui/material";
import { useFormik } from "formik";

const validationSchema = yup.object({
  username: yup.string("Enter your username").required("Username is required"),
  password: yup.string("Enter your password").required("Password is required"),
});

export const Login = () => {
  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
    },
    validationSchema,
    onSubmit: (values) => {
      login(values)
        .then((response) => {
          localStorage.setItem("Token", response.Jwt);
          toast.success("Successful login");
          //location.replace("/home");
        })
        .catch((err) => {
          console.log(err);
        });
    },
  });

  const fieldStyle = {
    paddingTop: "10px",
    paddingBottom: "10px",
    color: "primary",
  };

  const hasError = (fieldName) =>
    formik.touched[fieldName] && Boolean(formik.errors[fieldName]);
  const errorMessage = (fieldName) =>
    formik.touched[fieldName] && formik.errors[fieldName];

  return (
    <div className="base-container">
      <div className="header">Login</div>
      <div className="content">
        {/* <div className="image">
          <img src={loginImg} alt="" />
        </div> */}
        <div className="form">
          <form onSubmit={formik.handleSubmit}>
            <TextField
              sx={fieldStyle}
              type="text"
              name="username"
              fullWidth
              label="Username"
              value={formik.values.username}
              onChange={formik.handleChange}
              error={hasError("username")}
              helperText={errorMessage("username")}
            />
            <TextField
              sx={fieldStyle}
              type="password"
              name="password"
              fullWidth
              label="Password"
              value={formik.values.password}
              onChange={formik.handleChange}
              error={hasError("password")}
              helperText={errorMessage("password")}
            />
            <div className="footer">
              <button className="btn" type="submit" value="Submit">
                Login
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};
export default Login;

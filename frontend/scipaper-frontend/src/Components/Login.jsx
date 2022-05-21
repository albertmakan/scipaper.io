import React from "react";
import * as yup from "yup";
import { toast } from "react-toastify";
import { login } from "../Services/userService";
import { TextField } from "@mui/material";
import { useFormik } from "formik";
import Typography from "@mui/material/Typography";
import { Box } from "@mui/system";
import Button from "@mui/material/Button";
import { Link } from "react-router-dom";

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
          localStorage.setItem("Token", response);
          toast.success("Successful login");
          window.location.replace("/home");
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
    <Box
      sx={{
        marginLeft: 30,
        marginRight: 30,
        marginTop: 10,
        marginBottom: 0,
      }}
    >
      <Typography variant="h4" align="center">
        Login
      </Typography>
      <Typography component={Link} variant="h6" to="/register">
        Register
      </Typography>

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
        <Button variant="contained" type="submit">
          Login
        </Button>
      </form>
    </Box>
  );
};
export default Login;

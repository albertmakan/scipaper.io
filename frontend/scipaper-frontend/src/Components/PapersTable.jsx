import * as React from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { useNavigate } from "react-router-dom";
import { objectIdToDate, publish } from "../Services/paperService";
import { IconButton } from "@mui/material";
import { toast } from "react-toastify";

export default function PapersTable({ papers, isMy }) {
  const navigate = useNavigate();
  return (
    <TableContainer component={Paper}>
      <Table aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="right">AUTHOR</TableCell>
            <TableCell align="right">TITLE</TableCell>
            <TableCell align="right">
              {isMy ? "Created at" : "Published at"}
            </TableCell>
            {isMy && (
              <>
                <TableCell />
                <TableCell />
              </>
            )}
          </TableRow>
        </TableHead>
        <TableBody>
          {papers?.map((row, i) => (
            <TableRow
              key={i}
              sx={{
                "&:last-child td, &:last-child th": { border: 0 },
                cursor: "pointer",
              }}
            >
              <TableCell align="right">{row.author}</TableCell>
              <TableCell
                align="right"
                onClick={() =>
                  navigate(`/paper/${isMy ? row.id : row.paperId}`)
                }
              >
                {row.title}
              </TableCell>
              <TableCell align="right">
                {objectIdToDate(isMy ? row.id : row.paperId).toDateString()}
              </TableCell>
              {isMy && (
                <>
                  <TableCell align="right">
                    <IconButton
                      size="small"
                      onClick={() => {
                        navigate(`/edit/${row.id}`);
                      }}
                    >
                      edit
                    </IconButton>
                  </TableCell>
                  <TableCell align="right">
                    <IconButton
                      size="small"
                      onClick={() => {
                        publish(row.id).then(() => toast("Paper published"));
                      }}
                    >
                      pub
                    </IconButton>
                  </TableCell>
                </>
              )}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

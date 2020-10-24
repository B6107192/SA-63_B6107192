import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntClub } from '../../api/models/EntClub';
import DeleteIcon from '@material-ui/icons/Delete'
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [club, setClub] = useState<EntClub[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getClub = async () => {
     const res = await api.listClub({ limit: 10, offset: 0 });
     setLoading(false);
     setClub(res);
   };
   getClub();
 }, [loading]);
 
 const deleteClub = async (id: number) => {
   const res = await api.deleteClub({ id: id });
   setLoading(true);
 };
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">ชื่อชมรม</TableCell>
           <TableCell align="center">ประเภทชมรม</TableCell>
           <TableCell align="center">User</TableCell>
           <TableCell align="center">กิจกรรมชมรม</TableCell>
           <TableCell align="center">สถานที่</TableCell>
           <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         { club === undefined
         ? null
         : club.map((item) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.cLUBNAME}</TableCell>
             <TableCell align="center">{item.edges?.clubtypes?.cLUBETYPENAME}</TableCell>
             <TableCell align="center">{item.edges?.user?.uSERNAME}</TableCell>
             <TableCell align="center">{item.edges?.activity?.cLUBEACTIVITYNAME}</TableCell>
             <TableCell align="center">{item.edges?.location?.cLUBELOCATIONNAME}</TableCell>
             <TableCell align="center">
             <Button
                  onClick={() => {
                    if (item.id === undefined){
                      return;
                    }
                    deleteClub(item.id);
                  }}
                  style={{ marginLeft: 10 }}
                  variant="contained"
                  color="secondary"
                  startIcon={<DeleteIcon />}
                >
                  Delete
                </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}

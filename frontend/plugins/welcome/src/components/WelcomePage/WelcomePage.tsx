import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';

import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ระบบจัดการชมรม' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={` ${profile.givenName || 'to Backstage'}`}
       subtitle=""
       ><Link component={RouterLink} to="/savedata">
          <Button variant="contained" color="primary">
            เข้าระบบ
         </Button>
        </Link>
     </Header>
     <Content>
       <ContentHeader title="ข้อมูลชมรม">
         
       </ContentHeader>
       
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
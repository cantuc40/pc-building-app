import './App.css';
import {Routes, Route} from 'react-router-dom';

import HeaderBar from './components/headerbar';
import FooterBar from './components/footerbar';

import Main from './pages/main';
import NewPart from './pages/newpart';
import NewPC from './pages/newpc';
import Part from './pages/part';
import Parts from './pages/parts';
import PC from './pages/pc';
import PCS from './pages/pcs';
import partsDB from './pages/partsDB';



function App() {
  return (
    <div>
      <HeaderBar/>
      <Routes>
        <Route path='/' element={<Main/>} />
        <Route path='/pcs' element={<PCS/>} />
        <Route path='/pcs/:id' element={<PC/>} />
        <Route path='/pcs/newpc' element={<NewPC/>} />


        <Route path='/parts' element={<partsDB/>} />
        <Route path='/parts/newpart' element={<NewPart/>} />
        <Route path='/parts/:partname' element={<Parts/> } />
        <Route path='/parts/:partname/:id' element={<Part/> } />
        <Route path='/parts/:partname/new' element={<newSpecificPart/> } />


      </Routes>
      <FooterBar/>
    </div>
  );
}

export default App;

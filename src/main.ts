import './styles.css';
import { loadProjects } from './projects';
import { replaceCopyright } from './copyright';
import { inject } from '@vercel/analytics';

replaceCopyright();
inject();
loadProjects();

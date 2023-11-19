import { describe, expect, test } from 'vitest';
import { bggTestListItems, bggXmlTestListItems, cults3dTestListItems, githubTestListItems } from './test.data';
import { scrapeBgg, scrapeCults3d, scrapeGithub, upgradeBggImage } from '../src/projects';

describe('Projects', () => {
    test('scrape Github data', () => {
        const parser = new DOMParser();
        const githubMockDoc = parser.parseFromString(githubTestListItems, 'text/html');
        console.log(githubMockDoc.firstElementChild?.outerHTML);
        const projects = scrapeGithub(githubMockDoc);

        console.log(JSON.stringify(projects));

        expect(projects.length).toEqual(1);

        expect(projects[0].host).toEqual('github');
        expect(projects[0].title).toEqual('NDoolan360-Site');
        expect(projects[0].description).toEqual('My hand crafted personal website');
        expect(projects[0].url?.toString()).toEqual('https://github.com/NDoolan360/NDoolan360-Site');
        expect(projects[0].image?.highResSrc).toEqual('/images/github.png');
        expect(projects[0].image?.lowResSrc).toEqual(null);
        expect(projects[0].image?.alt).toEqual('Github Logo');
        expect(projects[0].programmingLanguage?.name).toEqual('TypeScript');
        expect(projects[0].programmingLanguage?.style).toEqual('background-color: #3178c6');
    });

    test('scrape Cults3D data', () => {
        const parser = new DOMParser();
        const cults3dMockDoc = parser.parseFromString(cults3dTestListItems, 'text/html');
        console.log(cults3dMockDoc.firstElementChild?.outerHTML);
        const projects = scrapeCults3d(cults3dMockDoc);

        console.log(JSON.stringify(projects));

        expect(projects.length).toEqual(2);

        expect(projects[0].host).toEqual('cults3d');
        expect(projects[0].title).toEqual('Reciprocating Rack and Pinion Fidget V2');
        expect(projects[0].description).toEqual(undefined);
        expect(projects[0].url?.toString()).toEqual(
            'https://cults3d.com/en/3d-model/gadget/reciprocating-rack-and-pinion-fidget-v2'
        );
        expect(projects[0].image?.highResSrc).toEqual('https://files.cults3d.com/{RRaP High-res Image Link}');
        expect(projects[0].image?.lowResSrc).toEqual(
            'https://images.cults3d.com/{RRaP Image Link}/https://files.cults3d.com/{RRaP High-res Image Link}'
        );
        expect(projects[0].image?.alt).toEqual('RRaPv2.png Reciprocating Rack and Pinion Fidget V2');
        expect(projects[0].programmingLanguage).toEqual(undefined);

        expect(projects[1].host).toEqual('cults3d');
        expect(projects[1].title).toEqual('Thought Processor');
        expect(projects[1].description).toEqual(undefined);
        expect(projects[1].url?.toString()).toEqual('https://cults3d.com/en/3d-model/art/thought-processor');
        expect(projects[1].image?.highResSrc).toEqual(
            'https://files.cults3d.com/{Thought Processor High-res Image Link}'
        );
        expect(projects[1].image?.lowResSrc).toEqual(
            'https://images.cults3d.com/{Thought Processor Image Link}/https://files.cults3d.com/{Thought Processor High-res Image Link}'
        );
        expect(projects[1].image?.alt).toEqual('Thought-Processor.png Thought Processor');
        expect(projects[1].programmingLanguage).toEqual(undefined);
    });

    test('scrape BGG data', () => {
        const parser = new DOMParser();
        const bggMockDoc = parser.parseFromString(bggTestListItems, 'text/html');
        console.log(bggMockDoc.firstElementChild?.outerHTML);
        const projects = scrapeBgg(bggMockDoc);

        expect(projects.length).toEqual(1);

        expect(projects[0].host).toEqual('boardgamegeek');
        expect(projects[0].title).toEqual('Cake Toppers');
        expect(projects[0].description).toEqual('Bakers assemble the most outrageous cakes to top each other.');
        expect(projects[0].url?.toString()).toEqual('https://boardgamegeek.com/boardgame/330653/cake-toppers');
        expect(projects[0].image?.highResSrc).toEqual('{Cake Toppers Image Link}');
        expect(projects[0].image?.lowResSrc).toEqual(null);
        expect(projects[0].image?.alt).toEqual('Board Game: Cake Toppers');
        expect(projects[0].programmingLanguage).toEqual(undefined);

        const bggMockXmlDoc = parser.parseFromString(bggXmlTestListItems, 'text/xml');
        upgradeBggImage(projects[0], bggMockXmlDoc);

        expect(projects[0].image?.highResSrc).toEqual('{Cake Toppers High-res Image Link}');
        expect(projects[0].image?.lowResSrc).toEqual('{Cake Toppers Image Link}');
    });
});

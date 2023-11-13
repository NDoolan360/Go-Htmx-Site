export type SectionValue = { type: string, value: string };

export default {
    "interests": [
        {
            "type": "chip",
            "value": "Web Development"
        },
        {
            "type": "chip",
            "value": "Coding Projects"
        },
        {
            "type": "chip",
            "value": "3D Printing"
        },
        {
            "type": "chip",
            "value": "Board Games"
        },
        {
            "type": "chip",
            "value": "Boardgame Design"
        },
        {
            "type": "chip",
            "value": "Tabletop Games"
        },
        {
            "type": "chip",
            "value": "Camping"
        },
        {
            "type": "chip",
            "value": "Hiking"
        },
        {
            "type": "chip",
            "value": "Cooking"
        }
    ],
    "projects": [
        {
            "type": "card",
            "value": "https://github.com/NDoolan360/Soundcheck"
        }
    ]
} as { [sectionName: string]: SectionValue[] };
import React from 'react';
import classes from './Cocktails.module.scss';
import Mojito from '../../data/resources/images/cocktails/mojito.png'
import Margarita from '../../data/resources/images/cocktails/margarita.jpg'
import AperolSpritz from '../../data/resources/images/cocktails/Aperol_Spritz.webp'
import BanananDaiquire from '../../data/resources/images/cocktails/Banana-Daiquiri.webp'
import MoscowMule from '../../data/resources/images/cocktails/moscow_mule.webp'
import EspressoMartini from '../../data/resources/images/cocktails/Espresso-Martini.png'
import Negroni from '../../data/resources/images/cocktails/Negroni.jpg'
import PinaColada from '../../data/resources/images/cocktails/pinacolada.webp'
import RumPunch from '../../data/resources/images/cocktails/rum_punch.jpg'
import TequilaSunrise from '../../data/resources/images/cocktails/tequila-sunrise.jpg'

interface Cocktail {
    id: number;
    name: string;
    description: string;
    ingredients: string[];
    image: string;
}

const cocktails: Cocktail[] = [
    {
        id: 1,
        name: "Mojito",
        description: "A refreshing Cuban highball that traditionally consists of five ingredients: white rum, sugar, lime juice, soda water, and mint.",
        ingredients: [
            "2 oz white rum",
            "1 oz fresh lime juice",
            "2 teaspoons sugar",
            "6-8 fresh mint leaves",
            "Soda water",
            "Lime wheel and mint sprig for garnish"
        ],
        image: Mojito
    },
    {
        id: 2,
        name: "Margarita",
        description: "A classic cocktail consisting of tequila, orange liqueur, and lime juice, often served with salt on the rim of the glass.",
        ingredients: [
            "2 oz tequila",
            "1 oz Cointreau or triple sec",
            "1 oz fresh lime juice",
            "Salt for rim",
            "Lime wheel for garnish"
        ],
        image: Margarita
    },
    {
        id: 3,
        name: "Espresso Martini",
        description: "A cold, coffee-flavored cocktail made with vodka, espresso coffee, and coffee liqueur.",
        ingredients: [
            "60 ml Vodka",
            "30 ml Fresh Espresso",
            "30 ml Coffee Liqueur",
            "15 ml Simple Syrup",
            "Ice"
        ],
        image: EspressoMartini
    },
    {
        id: 5,
        name: "Negroni",
        description: "A classic Italian cocktail that is considered an aperitif, made with gin, vermouth rosso, and Campari.",
        ingredients: [
            "30 ml Gin",
            "30 ml Campari",
            "30 ml Sweet Vermouth",
            "Orange Peel",
            "Ice"
        ],
        image: Negroni
    },
    {
        id: 6,
        name: "Piña Colada",
        description: "A sweet cocktail made with rum, coconut cream, and pineapple juice, usually served either blended or shaken with ice.",
        ingredients: [
            "2 oz white rum",
            "2 oz coconut cream",
            "2 oz pineapple juice",
            "Pineapple wedge and cherry for garnish"
        ],
        image: PinaColada
    },
    {
        id: 7,
        name: "Moscow Mule",
        description: "A cocktail made with vodka, spicy ginger beer, and lime juice, traditionally served in a copper mug.",
        ingredients: [
            "2 oz vodka",
            "4-6 oz ginger beer",
            "1/2 oz fresh lime juice",
            "Lime wheel for garnish"
        ],
        image: MoscowMule
    },
    {
        id: 9,
        name: "Rum Punch",
        description: "A tropical cocktail that combines dark rum with various fruit juices and grenadine for a sweet and fruity flavor.",
        ingredients: [
            "2 oz dark rum",
            "1 oz orange juice",
            "1 oz pineapple juice",
            "1/2 oz lime juice",
            "1/2 oz grenadine",
            "Orange slice and cherry for garnish"
        ],
        image: RumPunch
    },
    {
        id: 10,
        name: "Aperol Spritz",
        description: "A refreshing Italian aperitif cocktail made with prosecco, Aperol, and soda water.",
        ingredients: [
            "3 oz prosecco",
            "2 oz Aperol",
            "1 oz soda water",
            "Orange slice for garnish"
        ],
        image: AperolSpritz
    },
    {
        id: 11,
        name: "Tequila Sunrise",
        description: "A cocktail made with tequila, orange juice, and grenadine, creating a beautiful sunrise effect in the glass.",
        ingredients: [
            "2 oz tequila",
            "4 oz orange juice",
            "1/2 oz grenadine",
            "Orange slice and cherry for garnish"
        ],
        image: TequilaSunrise
    },
    {
        id: 12,
        name: "Banana Daiquiri",
        description: "A fruity variation of the classic daiquiri, made with rum, lime juice, simple syrup, and fresh banana.",
        ingredients: [
            "2 oz white rum",
            "1 oz fresh lime juice",
            "1/2 oz simple syrup",
            "1 ripe banana",
            "Banana slice for garnish"
        ],
        image: BanananDaiquire
    }
];

export const Cocktails: React.FC = () => {
    return (
        <div className={classes.cocktailsPage}>
            <div className={classes.header}>
                <h1>Cocktail Recipes</h1>
                <p>Discover and learn to make classic cocktails</p>
            </div>
            
            <div className={classes.cocktailsGrid}>
                {cocktails.map((cocktail) => (
                    <div key={cocktail.id} className={classes.cocktailCard}>
                        <div className={classes.imageContainer}>
                            <img src={cocktail.image} alt={cocktail.name} />
                        </div>
                        <div className={classes.content}>
                            <h2>{cocktail.name}</h2>
                            <p className={classes.description}>{cocktail.description}</p>
                            
                            <div className={classes.ingredients}>
                                <h3>Ingredients</h3>
                                <ul>
                                    {cocktail.ingredients.map((ingredient, index) => (
                                        <li key={index}>{ingredient}</li>
                                    ))}
                                </ul>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};
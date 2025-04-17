import React from 'react';
import classes from './Cocktails.module.scss';

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
            "2 oz White Rum",
            "1 oz Fresh Lime Juice",
            "2 tsp Sugar",
            "6-8 Fresh Mint Leaves",
            "Soda Water",
            "Ice"
        ],
        image: "https://images.unsplash.com/photo-1551538827-9c037cb4f32a?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"
    },
    {
        id: 2,
        name: "Old Fashioned",
        description: "A cocktail made by muddling sugar with bitters and water, adding whiskey or bourbon, and garnishing with orange slice or zest and a cocktail cherry.",
        ingredients: [
            "2 oz Bourbon or Rye Whiskey",
            "1 Sugar Cube",
            "2-3 dashes Angostura Bitters",
            "Orange Peel",
            "Ice"
        ],
        image: "https://images.unsplash.com/photo-1514362545857-3bc16c4c7d1b?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"
    },
    {
        id: 3,
        name: "Margarita",
        description: "A cocktail consisting of tequila, orange liqueur, and lime juice often served with salt on the rim of the glass.",
        ingredients: [
            "2 oz Tequila",
            "1 oz Cointreau or Triple Sec",
            "1 oz Fresh Lime Juice",
            "Salt for rim",
            "Ice"
        ],
        image: "https://images.unsplash.com/photo-1551538827-9c037cb4f32a?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"
    },
    {
        id: 4,
        name: "Espresso Martini",
        description: "A cold, coffee-flavored cocktail made with vodka, espresso coffee, and coffee liqueur.",
        ingredients: [
            "2 oz Vodka",
            "1 oz Fresh Espresso",
            "1 oz Coffee Liqueur",
            "1/2 oz Simple Syrup",
            "Ice"
        ],
        image: "https://images.unsplash.com/photo-1514362545857-3bc16c4c7d1b?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"
    },
    {
        id: 5,
        name: "Negroni",
        description: "A classic Italian cocktail that is considered an aperitif, made with gin, vermouth rosso, and Campari.",
        ingredients: [
            "1 oz Gin",
            "1 oz Campari",
            "1 oz Sweet Vermouth",
            "Orange Peel",
            "Ice"
        ],
        image: "https://images.unsplash.com/photo-1514362545857-3bc16c4c7d1b?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"
    },
    {
        id: 6,
        name: "Pina Colada",
        description: "A sweet cocktail made with rum, coconut cream, and pineapple juice, usually served either blended or shaken with ice.",
        ingredients: [
            "2 oz White Rum",
            "2 oz Coconut Cream",
            "3 oz Pineapple Juice",
            "Pineapple Wedge",
            "Ice"
        ],
        image: "https://images.unsplash.com/photo-1551538827-9c037cb4f32a?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"
    },
    {
        id: 7,
        name: "Moscow Mule",
        description: "A cocktail made with vodka, spicy ginger beer, and lime juice, garnished with a slice or wedge of lime.",
        ingredients: [
            "2 oz Vodka",
            "4 oz Ginger Beer",
            "1/2 oz Lime Juice",
            "Lime Wedge",
            "Ice"
        ],
        image: "https://images.unsplash.com/photo-1514362545857-3bc16c4c7d1b?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"
    },
    {
        id: 8,
        name: "Whiskey Sour",
        description: "A mixed drink containing whiskey, lemon juice, sugar, and optionally, a dash of egg white.",
        ingredients: [
            "2 oz Bourbon",
            "3/4 oz Fresh Lemon Juice",
            "3/4 oz Simple Syrup",
            "1/2 oz Egg White (optional)",
            "Ice"
        ],
        image: "https://images.unsplash.com/photo-1514362545857-3bc16c4c7d1b?ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"
    }
];

const Cocktails: React.FC = () => {
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

export default Cocktails;
//helpful sources: 
//https://docs.cypress.io/guides/core-concepts/introduction-To-Cypress
//https://docs.cypress.io/guides/end-to-end-testing/writing-your-first-end-to-end-test\
//https://docs.cypress.io/api/commands/get
//https://docs.cypress.io/guides/component-testing/angular/quickstart

describe('Test 1', () => {
  it('Creating a new review', () => {
    cy.visit('http://localhost:4200/');

    cy.get("input[id='location']", { timeout: 10000 }) // can also use cy.contains("Location") 
      .type("Lib West"); 
    cy.get("input[id='location']")
      .should('have.value', "Lib West"); 

    cy.get("input[id='rating']")
      .type("4"); 
    cy.get("input[id='rating']") 
      .should('have.value', "4"); 

    cy.get("input[id='description']") 
      .type("this is a teste"); 
    cy.get("input[id='description']") 
      .should('have.value',"this is a teste"); 

    cy.get("input[id='author']") 
      .type("Ya boi"); 
    cy.get("input[id='author']") 
      .should('have.value',"Ya boi"); 

    cy.contains("Submit Review") 
      .click(); 
    
  });
});

describe('Step 2', () => {
  it('Creating another review', () => {
    cy.visit('http://localhost:4200/');

    cy.get("input[id='location']", { timeout: 10000 }) // can also use cy.contains("Location") 
      .type("Carelton"); 
    cy.get("input[id='location']")
      .should('have.value', "Carelton"); 

    cy.get("input[id='rating']")
      .type("This is not a number"); 
    cy.get("input[id='rating']") 
      .should('have.value', "This is not a number"); 

    cy.get("input[id='description']") 
      .type("These chairs don't work"); 
    cy.get("input[id='description']") 
      .should('have.value',"These chairs don't work"); 

    cy.get("input[id='author']") 
      .type("Devala"); 
    cy.get("input[id='author']") 
      .should('have.value',"Devala"); 

    cy.contains("Submit Review") 
      .click(); 
    
  });
});

describe('Step 3', () => {
  it('Creating a third review', () => {
    cy.visit('http://localhost:4200/');

    cy.get("input[id='description']") 
      .type("this is a test3"); 
    cy.get("input[id='description']") 
      .should('have.value',"this is a test3"); 

    cy.get("input[id='author']") 
      .type("Emmet"); 
    cy.get("input[id='author']") 
      .should('have.value',"Emmet"); 

    cy.contains("Submit Review") 
      .click(); 
    
  });
});

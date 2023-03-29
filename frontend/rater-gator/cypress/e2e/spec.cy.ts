//helpful sources: 
//https://docs.cypress.io/guides/core-concepts/introduction-To-Cypress
//https://docs.cypress.io/guides/end-to-end-testing/writing-your-first-end-to-end-test\
//https://docs.cypress.io/api/commands/get
//https://docs.cypress.io/guides/component-testing/angular/quickstart

describe('Test 1', () => {
  it('Creating a new review', () => {
    cy.visit('http://localhost:3000/');

    cy.get("input[id='Name']", { timeout: 10000 }) // can also use cy.contains("Location") 
      .type("Turlington"); 
    cy.get("input[id='Name']")
      .should('have.value', "Turlington"); 

    cy.get("input[id='subject']", { timeout: 10000 }) // can also use cy.contains("Location") 
      .type("Turlington"); 
    cy.get("input[id='subject']")
      .should('have.value', "Turlington"); 

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

      Cypress.on('window:before:load', (win) => {
        cy.spy(win.console, 'error');
        cy.spy(win.console, 'warn');
      });
      
      
        cy.window().then((win) => {
          expect(win.console.error).to.have.callCount(0);
          expect(win.console.warn).to.have.callCount(0);
        });
  });
});

describe('Step 2', () => {
  it('Creating another review', () => {
    /* //Sprint 2 tests
    cy.visit('http://localhost:3000/');

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
      .click(); */

      cy.visit('http://localhost:3000/');

      cy.get("input[id='Name']", { timeout: 10000 }) // can also use cy.contains("Location") 
        .type("Turlington"); 
      cy.get("input[id='Name']")
        .should('have.value', "Turlington"); 
  
      cy.get("input[id='subject']", { timeout: 10000 }) // can also use cy.contains("Location") 
        .type("Carleton"); 
      cy.get("input[id='subject']")
        .should('have.value', "Carleton"); 
  
      cy.get("input[id='rating']")
        .type("2"); 
      cy.get("input[id='rating']") 
        .should('have.value', "2"); 
  
      cy.get("input[id='description']") 
        .type("this is a testeeee"); 
      cy.get("input[id='description']") 
        .should('have.value',"this is a testeeee"); 
  
      cy.get("input[id='author']") 
        .type("hey"); 
      cy.get("input[id='author']") 
        .should('have.value',"hey"); 
  
      cy.contains("Submit Review") 
        .click(); 
    
        
        
          cy.window().then((win) => {
            expect(win.console.error).to.have.callCount(1);
            expect(win.console.warn).to.have.callCount(0);
          });
    
  });
});

describe('Step 3', () => {
  it('Creating a third review', () => {

    /* //Sprint 2 tests
    cy.visit('http://localhost:3000/');

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
      */
    
      cy.visit('http://localhost:3000/');
      cy.contains("Submit Review") 
      .click(); 
    
      
      
        cy.window().then((win) => {
          expect(win.console.error).to.have.callCount(0);
          expect(win.console.warn).to.have.callCount(0);
        });

  });
});

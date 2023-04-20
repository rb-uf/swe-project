//helpful sources: 
//https://docs.cypress.io/guides/core-concepts/introduction-To-Cypress
//https://docs.cypress.io/guides/end-to-end-testing/writing-your-first-end-to-end-test\
//https://docs.cypress.io/api/commands/get
//https://docs.cypress.io/guides/component-testing/angular/quickstart

describe('Test 1', () => {
  it('Adding a new location', () => {
    cy.visit('http://localhost:3000/');

    cy.get("mat-select[id='subjectSelect']").click().wait(3000).focus().type('{esc}');

    cy.get("input[id='Name']", { timeout: 10000 }) // can also use cy.contains("Location") 
      .type("Turlington"); 
      cy.contains("Add Location") 
      .click();
      
    cy.get("mat-select[id='subjectSelect']").click().wait(5000);
  });
});

describe('Step 2', () => {
  it('adding a review', () => {

      cy.visit('http://localhost:3000/');

      cy.get('button[class*="example-icon"]').click();
  
      cy.get("input[id='subject']", { timeout: 10000 }) // can also use cy.contains("Location") 
        .type("Turlington"); 
      cy.get("input[id='subject']")
        .should('have.value', "Turlington"); 
  
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

        cy.wait(3000);
  
      cy.contains("Submit Review") 
        .click();
      
        cy.contains("Rater Gator") 
        .click();

        cy.wait(3000);
      
      cy.get("mat-select[id='subjectSelect']").click().get('mat-option').contains('Turlington').click();

      cy.wait(5000);
    
  });
});

describe('Step 3', () => {
  it('adding a second review', () => {

      cy.visit('http://localhost:3000/');

      cy.get('button[class*="example-icon"]').click();
  
      cy.get("input[id='subject']", { timeout: 10000 }) // can also use cy.contains("Location") 
        .type("Turlington"); 
      cy.get("input[id='subject']")
        .should('have.value', "Turlington"); 
  
      cy.get("input[id='rating']")
        .type("3"); 
      cy.get("input[id='rating']") 
        .should('have.value', "3"); 
  
      cy.get("input[id='description']") 
        .type("this is another testeeee"); 
      cy.get("input[id='description']") 
        .should('have.value',"this is another testeeee"); 
  
      cy.get("input[id='author']") 
        .type("Devala"); 
      cy.get("input[id='author']") 
        .should('have.value',"Devala"); 

        cy.wait(3000);
  
      cy.contains("Submit Review") 
        .click();
      
        cy.contains("Rater Gator") 
        .click();

        cy.wait(3000);
      
      cy.get("mat-select[id='subjectSelect']").click().get('mat-option').contains('Turlington').click();

      cy.wait(5000);
    
  });
});

describe('Step 4', () => {
  it('Deleting first review', () => {
    
      cy.visit('http://localhost:3000/');
      cy.get("mat-select[id='subjectSelect']").click().get('mat-option').contains('Turlington').click();
      cy.contains("Delete").click(); 
      cy.wait(5000);

      cy.reload().wait(1000);
      cy.get("mat-select[id='subjectSelect']").click().get('mat-option').contains('Turlington').click();
      cy.wait(5000);
  });
});

describe('Step 5', () => {
  it('Deleting a review', () => {
    
      cy.visit('http://localhost:3000/');
      cy.get("mat-select[id='subjectSelect']").click().get('mat-option').contains('Turlington').click();
      cy.contains("Delete").click(); 
      cy.wait(5000);

      cy.reload().wait(1000);
      cy.get("mat-select[id='subjectSelect']").click().get('mat-option').contains('Turlington').click();
      cy.wait(5000);
  });
});
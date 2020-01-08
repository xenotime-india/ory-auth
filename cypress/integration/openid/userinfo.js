import { prng } from '../../helpers';

describe('OpenID Connect Userinfo', () => {
  const nc = () => ({
    client_id: prng(),
    client_secret: prng(),
    scope: 'openid',
    redirect_uris: [`${Cypress.env('client_url')}/openid/callback`],
    grant_types: ['authorization_code', 'refresh_token']
  });

  it('should return a proper userinfo response', function() {
    const client = nc();
    cy.authCodeFlow(client, { consent: { scope: ['openid'] } }, 'openid');

    cy.get('body')
      .invoke('text')
      .then(content => {
        const { result } = JSON.parse(content);
        expect(result).to.equal('success');
      });

    cy.request(`${Cypress.env('client_url')}/openid/userinfo`)
      .its('body')
      .then(({ sub, sid } = {}) => {
        expect(sub).to.eq('foo@bar.com');
        expect(sid).to.not.be.empty;
      });
  });
});

import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FusionMainComponent } from './fusion-main.component';

describe('FusionMainComponent', () => {
  let component: FusionMainComponent;
  let fixture: ComponentFixture<FusionMainComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FusionMainComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FusionMainComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
